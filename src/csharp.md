

# build from cli

打开 visual studio prompt


```
csc a.cs
```

```cs
using System;
namespace HelloWorldApplication
{
   class HelloWorld
   {
      static void Main(string[] args)
      {
         Console.WriteLine("Hello World");
         Console.ReadKey();
      }
   }
}
```



# fmt format string

```c#
string.Format("blabla: {0}", 123);
string.Format("blabla: {0:d}", 123);
```


# switch

```c#
switch (somevar) 
{
	case 123:
		// do something
		// do more, no need bracket
		break;
	case 456:
	case 457:
		break;
	default:
		break;
}
```

# file

```csharp
// read file
string content = File.ReadAllText(path, Encoding.UTF8);

// write file
File.WriteAllText(curFile, "blabla");
```


