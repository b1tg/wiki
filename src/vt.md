# vt


# links

related repos:
- [SimpleVisor](https://github.com/ionescu007/SimpleVisor) VT-x hypervisor works on Windows and UEFI by ionescu007, “中心节点”之一, 
    - [源码阅读](./SimpleVisorRead.html)
- [HyperPlatform](https://github.com/tandasat/HyperPlatform) VM-exit filtering platform ；“中心节点” 之一
- [kHypervisor](https://github.com/Kelvinhack/kHypervisor) capable for nested virtualization in Windows x64 platform, 基于 HyperPlatform；
- [HypervisorKeylogger](https://github.com/ParkHanbum/HypervisorKeylogger) korean; 单文件vt、在xp上测试成功；
- [HackSysExtremeVulnerableDriver](https://github.com/hacksysteam/HackSysExtremeVulnerableDriver) Vulnerable Driver; 提供编译好的sys, .sln编译有问题
- [gbhv](https://github.com/Gbps/gbhv) 专注于 EPT hooking; README值得看； *#checkme*
- [chocolate_milk: Pure Rust x86_64 bootloader and kernel](https://github.com/gamozolabs/chocolate_milk) 常看常新
- [barbervisor](https://blog.talosintelligence.com/2020/08/barbervisor.html) for fuzz
- [HyperDbg](https://github.com/HyperDbg/HyperDbg) for debug | 貌似和 hypervisor-from-scratch 系列文章有关
- [VT_Learn](https://github.com/zzhouhe/VT_Learn) VT技术入门; 在xp上测试成功

tutorials:

- hypervisor development series:
    - [https://revers.engineering/day-0-virtual-environment-setup-scripts-and-windbg/](https://revers.engineering/day-0-virtual-environment-setup-scripts-and-windbg/)
    - [https://revers.engineering/day-1-introduction-to-virtualization/](https://revers.engineering/day-1-introduction-to-virtualization/)
    - [https://revers.engineering/day-2-entering-vmx-operation/](https://revers.engineering/day-2-entering-vmx-operation/)
    - [https://revers.engineering/day-3-multiprocessor-initialization-error-handling-the-vmcs/](https://revers.engineering/day-3-multiprocessor-initialization-error-handling-the-vmcs/)
    - [https://revers.engineering/day-4-vmcs-segmentation-ops](https://revers.engineering/day-4-vmcs-segmentation-ops)
    - [https://revers.engineering/day-5-vmexits-interrupts-cpuid-emulation/](https://revers.engineering/day-5-vmexits-interrupts-cpuid-emulation/)


- hypervisor from scratch series:
    - [Part 1: Basic Concepts & Configure Testing Environment](https://rayanfam.com/topics/hypervisor-from-scratch-part-1/)
    - [Part 2: Entering VMX Operation](https://rayanfam.com/topics/hypervisor-from-scratch-part-2/)
    - [Part 3: Setting up Our First Virtual Machine](https://rayanfam.com/topics/hypervisor-from-scratch-part-3/)
    - [Part 4: Address Translation Using Extended Page Table (EPT)](https://rayanfam.com/topics/hypervisor-from-scratch-part-4/)
    - [Part 5: Setting up VMCS & Running Guest Code](https://rayanfam.com/topics/hypervisor-from-scratch-part-5/)
    - [Part 6: Virtualizing An Already Running System](https://rayanfam.com/topics/hypervisor-from-scratch-part-6/)
    - [Part 7: Using EPT & Page-Level Monitoring Features](https://rayanfam.com/topics/hypervisor-from-scratch-part-7/)
    - [Part 8: How To Do Magic With Hypervisor!](https://rayanfam.com/topics/hypervisor-from-scratch-part-8/)

- VT技术入门 [https://space.bilibili.com/37877654/](https://space.bilibili.com/37877654/)
    - [VT技术入门01_资源和基础](https://www.bilibili.com/video/BV1hb411i7Hj)
    - [VT技术入门02_混合编译配置](https://www.bilibili.com/video/BV1cb411e7oU)
    - [VT技术入门03_支持检测](https://www.bilibili.com/video/BV1Eb411e7U9)
    - [VT技术入门04_VMXON](https://www.bilibili.com/video/BV1vb411e7LS)
    - [VT技术入门05_VMCS(1)](https://www.bilibili.com/video/BV1Pb411v78s)
    - [VT技术入门06_VMCS(2)](https://www.bilibili.com/video/BV1ob411n7jr)
    - [VT技术入门07_VMCS(3)](https://www.bilibili.com/video/BV1fb411n7GT)
    - [VT技术入门08_调试技巧](https://www.bilibili.com/video/BV1Gb411H7UW)
    - [VT技术入门09_VM-Exit_handler](https://www.bilibili.com/video/BV1mb411H7LX)
    - [VT技术入门10_EPT物理地址转换](https://www.bilibili.com/video/BV1fb411H73R)
    - [VT技术入门11_非PAE下的EPT开启](https://www.bilibili.com/video/BV1mb411J7Wu)
    - [VT技术入门12_PAE下的EPT开启](https://www.bilibili.com/video/BV15b411n72U)
    - [VT技术入门13_应用举例](https://www.bilibili.com/video/BV1Hb411n7Mw)