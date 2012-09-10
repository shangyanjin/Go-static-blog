title: Windows'da gnu/linux tadı – Cygwin
datetime: 2011/08/05 17:01
language: turkish
description: Cygwin Windows işletim sistemlerinde, Linux benzeri bir ortam sağlar.

Cygwin isimli program sayesinde gnu/linux severler windows platformunda kendini
yabancı hissetmeyecek. Bu yazıda kısaca Windows ortamında gnu/linux görünüm ve
hissini sağlayan araçlar bütünü olan cygwin’den bahsedeceğim.

Cygwin isimli araçlar bütününü yeni keşfettim ve bir hayli de beğendim. Cygwin
sayesinde windows altında bash, wget, python, rsync, openssh ve gnu/linux platformlarda
kullanmaya alıştığımız daha nice programı kolaylıkla kurup kullanabiliyoruz.

Cgywin’in kurulumu oldukça basit. [Buradan](http://www.cygwin.com/setup.exe) indirdirip
çalıştırdıktan sonra, kurulumun ilk safhasında kurulumu nereden yapacağımızı soruyor, bu
aşamada çoğu durumda "Install from internet" (Internetten kur) seçeneğini seçmeniz gerekiyor.
Daha sonra size, indirilen dosyaların nerede tutulacağını, ve cygwin’in kök dizininin neresi
olacağını soruyor. Cygwin’in kök dizini, linux sistemlerdeki kök dizinin görevini görecek.
Nereye yüklediğiniz çok fark etmeyecek, öntanımlı olarak `C:\cgywin` içerisine kuruyor.
Kurulumu tamamladıktan sonra bu dizinin içinde, `home`, `lib`, `bin` gibi dizinler içinde,
cgywin ile birlikte kurduğunuz programların dosyalarını bulabilirsiniz. Ama doğrudan bu
dizindeki dosyaları kullanmayacaksınız. O kısıma geleceğiz. Dosya yollarını seçtikten sonra
sizden bir url seçmenizi isteyecek. Bu linki Cgywin ile birlikte kullanılacak araçları seçmek
için kullanacak. Herhangi birini seçebileceğiniz tahmin ediyorum.

Daha sonraki ekranda sizden program seçmenizi isteyecek, burada, kategorilerin altından istediğiniz
programları seçip kuruluma devam edebilirsiniz. En son ekrandan masaüstüne kısayol eklemeyi unutmayın,
gerekli olacak.

Kurulum bittikten sonra, masaüstüne veya programlar menüsüne eklediğiniz cgywin ikonuna tıklayarak
cgywin kabuğuna erişebilirsiniz. Cgywin kabuk ile kurduğunuz bütün programlara bu kabuk üzerinden
ulaşabileceksiniz.

Benim bunu yazdığım tarihde python kategorisi içerisinde setuptools yoktu. Setuptools yüklemek için
cgywin kabuk içerisinde wget ile [buradaki dosyayı](http://peak.telecommunity.com/dist/ez_setup.py)
indirdikten sonra, python ile çalıştırmanız gerekecek. Indirdiğiniz python betiği setuptools için
gerekli egg dosyasını indirip kurma işini kendi halledecek. Bu aşamadan sonra [Python paket indeksindeki](http://pypi.python.org/pypi?%3Aaction=browse)
herhangi bir paketi indirip, setuptools’un içinde gelen easy_install ile rahatça kurabilirsiniz.