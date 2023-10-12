# expiration
Small command-line utility to retrieve the running Windows build's expiration date. It works by retrieving the SystemExpirationDate value from [KUSER_SHARED_DATA](https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/ntddk/ns-ntddk-kuser_shared_data).

![image](https://github.com/dgurney/expiration/assets/12816807/c4a05e88-30b9-4f45-9e33-b8ea55ac9fb5)

## Compatibility 
It has been tested to work on 64-bit (including arm64) Windows versions as low as 7.

## Thanks
https://github.com/dhrdlicka/timebomb gave me the needed memory address on a silver platter, saving a bit of time from needing to find it in the `winver.exe` disassembly.
