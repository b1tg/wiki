# Source Code Read: SimpleVisor  [Pinback: vt](./vt.html)




从驱动入口 DriverEntry 开始读：

首先做了一个 power state callback，之后加载 hypervisor

hypervisor 加载的逻辑在函数 ShvLoad 中，剩下的阅读都从这个函数出发。


首先需要让所有的 LP（Logic Processor） 进入 VMX root 模式


```c++
ShvLoad
	ShvOsRunCallbackOnProcessors(ShvVpLoadCallback, &callbackContext);
		KeGenericCallDpc // callback
			ShvOsDpcRoutine // callback
				ShvVpLoadCallback // 在这个函数中cpu 被 hyperjacked，~返回后就是unload和清理了~
				// do cleanup:
				// ShvVmxCleanup, KeSignalCallDpcSynchronize, KeSignalCallDpcDone
				
				


ShvVpLoadCallback
	ShvVmxProbe // 检查 VMX root mode 是否支持
	PSHV_VP_DATA vpData // 初始化vpData，代表一个~LP~的数据结构
	vpData->SystemDirectoryTableBase = Context->Cr3; // TODO: 看注释，什么是PML4
	ShvVpInitialize(vpData) // 初始化VP
	if (ShvIsOurHypervisorPresent() == FALSE) // 通过cpuid检查当前hypervisor是否初始化成功
											  // ~此时cpuid应该已经被hook~		
	InterlockedIncrement(Context->InitCount); // 增加初始化成功的LP计数， This CPU is hyperjacked!



ShvVpInitialize(Data)
	ShvOsPrepareProcessor // dump,只在uefi时有用
	ShvCaptureSpecialRegisters(&Data->SpecialRegisters); // 初始化vpData的SpecialRegisters:
														  // cr0, cr3, cr4, debug_ctl, msr_gs_base,
														  // kernel_dr7, gdtr.limit, idtr.limit, tr, ldtr
	ShvOsCaptureContext(&Data->ContextFrame);
		RtlCaptureContext(ContextRecord); // windows提供的函数：
										  // Retrieves a context record in the context of the caller.
	status = ShvVmxLaunchOnVp(Data);  // 如果EFLAGS_ALIGN_CHECK没有set，为当前处理器初始化VMX
		

ShvVmxLaunchOnVp(VpData)
	VpData->MsrData[i].QuadPart = readmsr(MSR_IA32_VMX_BASIC + i); // 初始化VMX相关的一系列MSRs
	ShvVmxMtrrInitialize(VpData);  // 初始化 MTRR (Memory Type Range Registers) 相关的一系列MSRs，
								   // 这似乎是为了EPT准备的。				
	ShvVmxEptInitialize // 初始化 EPT structures
	ShvVmxEnterRootModeOnVp(VpData) // Attempt to enter VMX root mode on this processor.
	ShvVmxSetupVmcsForVp(VpData); // 初始化 VMCS，包括guest和host state，大多数“可选监控项”都在这里面
	ShvVmxLaunch()
		vmlaunch // 这里不该返回了，返回就是失败了，后面是失败处理
				 // 如果成功rip会指向 ShvVpRestoreAfterLaunch
		failureCode = (INT32)ShvVmxRead(VM_INSTRUCTION_ERROR);
		vmoff 
	
ShvVmxEnterRootModeOnVp(VpData)
	// - 做一些 check
	// - 写 revision ID到VmxOn和vmcs里面
	// - 设置一些physical addresses
	// - 奇怪的掩码方式设置CR0和cr4
	__vmx_on // Enable VMX Root Mode
	__vmx_vmclear // Clear the state of the VMCS, setting it to Inactive
	__vmx_vmptrld // Load the VMCS, setting its state to Active
	// 至此 VMX Root Mode is enabled, with an active VMCS.
	
==== 重要的数据结构 ====

typedef struct _SHV_VP_DATA
{
    union
    {
        DECLSPEC_ALIGN(PAGE_SIZE) UINT8 ShvStackLimit[KERNEL_STACK_SIZE];
        struct
        {
            SHV_SPECIAL_REGISTERS SpecialRegisters;
            CONTEXT ContextFrame;
            UINT64 SystemDirectoryTableBase;
            LARGE_INTEGER MsrData[17];
            SHV_MTRR_RANGE MtrrData[16];
            UINT64 VmxOnPhysicalAddress;
            UINT64 VmcsPhysicalAddress;
            UINT64 MsrBitmapPhysicalAddress;
            UINT64 EptPml4PhysicalAddress;
            UINT32 EptControls;
        };
    };

    DECLSPEC_ALIGN(PAGE_SIZE) UINT8 MsrBitmap[PAGE_SIZE];
    DECLSPEC_ALIGN(PAGE_SIZE) VMX_EPML4E Epml4[PML4E_ENTRY_COUNT];
    DECLSPEC_ALIGN(PAGE_SIZE) VMX_PDPTE Epdpt[PDPTE_ENTRY_COUNT];
    DECLSPEC_ALIGN(PAGE_SIZE) VMX_LARGE_PDE Epde[PDPTE_ENTRY_COUNT][PDE_ENTRY_COUNT];

    DECLSPEC_ALIGN(PAGE_SIZE) VMX_VMCS VmxOn;
    DECLSPEC_ALIGN(PAGE_SIZE) VMX_VMCS Vmcs;
} SHV_VP_DATA, *PSHV_VP_DATA;
```
