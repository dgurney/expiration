# expiration
Small command-line utility to retrieve the running Windows build's expiration date. It works by retrieving the SystemExpirationDate value from [KUSER_SHARED_DATA](https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/ntddk/ns-ntddk-kuser_shared_data).

![image](https://github.com/dgurney/expiration/assets/12816807/f465f804-beb4-4545-8f7d-0f787e86198b)

## Compatibility 
It has been tested to work on Windows 10/11, and Windows 7 (with latest updates).

## Thanks
https://github.com/dhrdlicka/timebomb gave me the needed memory address on a silver platter, saving a bit of time from needing to find it in the `winver.exe` disassembly.
