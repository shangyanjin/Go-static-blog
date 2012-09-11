title: Django şablonlarında php
datetime: 2011/09/06 12:44
language: turkish

django_php paketiyle django uygulamalarınızın web şablonlarında php kullanmanız
mümkün. Bunu neden yapmak isteyeceğiniz ise tamamen bir muamma…

djang_php’yi

    easy_install django_php

veya

    pip install django_php

komutlarıyla kurabilirsiniz. Tabi ki bu komutları vermek için gerekli python
paketlerinin yüklü olmasını gerektiğini unutmamak gerek.

Daha sonra django_php’yi settings modülü içerisindeki `INSTALLED_APPS` listesine
ekleyerek kullanmaya başlayabilirsiniz.

{% codeblock lang:python %}
INSTALLED_APPS = (
   'django_php',
)
{% endcodeblock %}

django_php’yi kullanmadan önce php_cgi’in sisteminizde yüklü olduğundan emin olun.
Eğer php_cgi’in yerini belirtmeniz gerekiyorsa, settings modülü içerisinde

{% codeblock lang:python %}
PHP_CGI = '/usr/local/bin/php-cgi'
{% endcodeblock %}

şeklinde belirtebilirsiniz. Çoğu zaman bu ayarı yapmanıza gerek yoktur.

Şablonlarınızın içine php’yi dahil etmek için de, şablon dosyalarınızın içinde:

    {% load php %}
    {% php echo 9; %}

şeklinde kullanabilirsiniz. Daha fazla örnek kaynak kodlarının içinde mevcut.

İlgili Linkler:

 - [Proje anasayfası](http://animuchan.net/django_php/)
 - [PyPi Sayfası](http://pypi.python.org/pypi/django_php)
 - [Kodlar](https://github.com/mvasilkov/django-php)