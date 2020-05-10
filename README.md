# Web Attack Log Simulation

## Description
A basic program made in Go to simulate a web attack via logs. I did this to test some Linux commands I could use for when I am getting attacked by thousands of malicious sources (DDoS) and it's hammering PHP FPM.

## Program's Parameters/Arguments

* 1 - The file path to append lines to.

* 2 - The amount of lines to write up to until the program closes.

* 3 - The amount of time between each write (in milliseconds).

Example:

```
/home/roy/go_build_fakeAttack_linux /var/log/fake_attack.log 90000 50
```

## Credits
* Christian Deacon (AKA Roy) - Writing the entire application.