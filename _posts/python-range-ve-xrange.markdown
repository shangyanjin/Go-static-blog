title: Python range ve xrange
language: turkish
datetime: 2011/08/02 13:22
description: Python programlama dilindeki range ve xrange fonksiyonu arasında ne tür faklılıklar vardır? Python 2.x ve 3.x arasında nasıl uyumlu olarak çalıştırılır.

Python 2 ile python 3 arasında `range` fonksiyonu farklılık gösteriyor. Python
betiklerinde kullanılan bu fonksiyon, eğer doğru python yorumlayıcısında
çalıştırılmazsa, istenildiğinden farklı davranabilir. Bu sorundan kurtulmak için,
aşağıdaki yöntemi kullanıyorum.

Yöntemden bahsetmeden önce, sorun hakkında biraz bilgi vereceğim. Python 2
sürümünde, `range` ve `xrange` adıyla iki farklı fonksiyon var. `range` isimli fonksiyon,
bir liste döndürüyor. `xrange` isimli fonksiyon ise bir "generator" (tr: üretici) fonksiyon.
Bu iki fonksiyon arasındaki fark, hafıza kullanımında. `xrange` fonksiyonu her çağırıldığında
yeni bir obje döndürdüğü için, daha az hafıza kullanılıyor.

`range` ve `xrange` arasındaki bu fark nedeniyle, programlarınızda `xrange` fonksiyonunu tercih
edenlerdenseniz, kodlarınızı python 3 yorumlayıcı çalışıtırmayacaktır. Çünkü python 3 ile birlikte,
`xrange` fonksiyonu kaldırıldı ve `range` fonksiyonu, python 2'deki `xrange` fonksiyonu gibi davranmaya başladı.

Aşağıda görülebilen örnek kod ile, python sürümleri arasındaki farkdan oluşan bu sorunun üstesinden gelebilirsiniz.
Bu kodları modülünüzün yukarılarında kullanmalı, ve `xrange` kullanmak yerine range kullanmayı tercih etmelisiniz.
Bu kodun çalıştığı platforma göre, `xrange` ve `range` fonksiyonu kendiliğinden kullanılacak.

{% codeblock lang:py %}
from sys import version_info
if version_info[0] == 2:
    range = xrange
{% endcodeblock %}