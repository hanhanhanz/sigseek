# Sigseek


simple script for finding strings to fingerprinting anti-root or ssl-pinning method for recon. because i am too lazy for searching every keyword in [JADX](https://github.com/skylot/jadx)


### Requirement

[Apktool](https://github.com/iBotPeaches/Apktool)

### Installation

```sh
$ go build sigseek.go
$ ./sigseek
```


### Example
```sh
$ ./sigseek -f UnCrackable-Level1.apk
/dev | ss-UnCrackable-Level1.apk/smali/sg/vantagepoint/a/c.smali | 97
/system/xbin/ | ss-UnCrackable-Level1.apk/smali/sg/vantagepoint/a/c.smali | 87
Superuser.apk | ss-UnCrackable-Level1.apk/smali/sg/vantagepoint/a/c.smali | 85
/system/bin/.ext/ | ss-UnCrackable-Level1.apk/smali/sg/vantagepoint/a/c.smali | 91
/system/app/Superuser.apk | ss-UnCrackable-Level1.apk/smali/sg/vantagepoint/a/c.smali | 85
com.koushikdutta.superuser | ss-UnCrackable-Level1.apk/smali/sg/vantagepoint/a/c.smali | 97
/system/bin/ | ss-UnCrackable-Level1.apk/smali/sg/vantagepoint/a/c.smali | 91

```

