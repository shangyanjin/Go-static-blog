title: Django'da pdb ile debug
datetime: 2011/08/04 11:41
language: turkish
description: Bu makalede Django uygulamalarını pdb kullanarak nasıl hatalardan ayıklayabileceğimizden bahsediyor.

Django’da geliştirdiğimiz web uygulamasının hata temizlemesini isterseniz pdb
(python debugger) ile de yapabilirsiniz. Bu yazıda kısaca bunun nasıl
yapıldığından bahsedeceğiz.

Django’nun kendine ait bir debug aracı var, ama django ile python debugger
kullanmak isteyenler için, django-pdb var. django-pdb sayesinde django
uygulamalarımızı pdb ile debug edebiliriz. django-pdb’nin kurulumu `pip`
ile kolayca yapılabilir. `pip install django-pdb` komutu django-pdb’nin
kurulumunu sizin için yapacaktır. *Nix kullananların kendi dağıtımlarına
ait depoları kontrol etmelerinde de fayda var. Eğer depolarda bulabiliyorsanız,
kendi paket yöneticinizle de kurabilirsiniz.
{% codeblock lang:bash %}
pip install django-pdb
{% endcodeblock %}

Kurulumu tamamladıktan sonra django ayar dosyanızdaki, yüklü uygulamalara
(INSTALLED_APPS) django_pdb’yi ekleyerek django’da geliştirdiğimiz siteye dahil
ediyoruz.

{% codeblock lang:python %}
INSTALLED_APPS = (
  ’django_pdb’,
)
{% endcodeblock %}
Debugger’ın çalışması için birkaç farklı yöntem var, ama hepsi için settings
modülündeki `DEBUG` değişkeninin, `True`’ya eşitlenmesi gerekiyor. Aksi halde
çalışmayacaktır. `settings.DEBUG`’ın `True` olduğundan emin olduktan sonra,
ek bir işlem yapmadan django’nun kendi geliştirme sunucusunu başlatabilirsiniz.
GET metodunda pdb olan herhangi bir sayfa’yı açmaya çalıştığınızda pdb devreye
girecektir. (ÖRN: www.ornek.com/?pdb)

{% codeblock lang:python %}
DEBUG = True
{% endcodeblock %}

Eğer geliştirme sunucunuzu `--pdb` anahtarı ile başlatırsanız, yüklediğiniz her
view sayfasıyla birlikte pdb devreye girecektir.

{% codeblock lang:python %}
manage.py runserver --pdb
{% endcodeblock %}

django-pdb’nin python paket indeksi (pypi) sayfasına da [buradan](http://pypi.python.org/pypi/django-pdb) ulaşabilirsiniz.