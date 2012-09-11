title: Django ile Blog Geliştirme - İlk ayarlar
datetime: 2011/09/05 11:14
description: Django web çatısında blog geliştirme yazı dizisi, 1. bölüm. Bu yazıda konuya başlangıç yapacağız.
language: turkish

Python ile web geliştirmenin benim için vazgeçilmezlerinden biri olan django
çatısı ile, baştan başlayarak nasıl bir blog geliştirileceğine dair bir yazı
dizisine başlıyorum. 0'dan blog'a django yazı dizisinin 1. bölümünde yeni bir
django projesine nasıl başlanacağına ve ilk ayarlara değineceğim.

Bu yazı dizisinde bahsedilen blog uygulamasının son haline [github deposundan](https://github.com/yasar11732/django-blog) ulaşabilirsiniz.

Django çatısı ile python web uygulaması geliştirmeye başlamadan önce ilk iş
bir Python yorumlayıcısı edinmek olacaktır. Eğer bir gnu/linux veya unix tabanlı
bir işletim sistemi kullanıyorsanız, muhtemelen Python programlama dili
sisteminizde zaten kuruludur. Emin olmak için, gnu/linux komut satırına `which Python`
komutunu verebilirsiniz. Eğer Python sisteminizde yüklü değilse, komutun
bulunamadığına dair bir hata alırsınız. Böyle bir hata almanız halinde, sisteminize
ait paket yöneticisi aracılığıyla, veya kaynak koddan derleyerek Python programlama
dilini kurabilirsiniz. Bunların detaylarına değinmeyeceğim.

Ancak şunu belirtmek isterim ki Django Python'un 3.0 ve üzeri sürümleri desteklemiyor. Python
kullanmak için, 2.5,2.6 veya 2.7 sürümlerini tavsiye edebilirim.

Windows'da Python kurmak için ise [Python'un resmi sitesi](http://python.org)nden Windows kurulum
dosyalarını indirip, kurmanız gerekiyor. 2.x bir sürüme ihtiyacınız olduğunu
tekrar belirtmek istiyorum. Python'u kurduktan sonra, Python ana dizinini
(genellikle `C:\Python27`) ve Python ana dizini içindeki `Scripts` dizinini
windows sistem yoluna eklemek iyi bir fikir olabilir. Böylece, Python'un kendisi
ve yüklü Python paketlerinin çalıştırılabilir dosyaları komut satırından erişilebilir
hale gelir. Windows'da Python kullanmanın bir diğer yöntemi için [Cygwin aracıyla
ilgili açıklamalar]({% relative /blog/turkish/2011/08/05/cygwin-ozellikleri.html %})'a
bakabilirsiniz.

Çalışan bir Python yorumlayıcısı elde ettikten sonra, ihtiyacınız olan şey -tahmin
edersiniz ki- django çatısı. Gnu/Linux sistemlerin birçoğunda Django web geliştirme
çatısı resmi paket depolarında veya kullanıcı depolarında bulunabilir. Eğer
depolarda bulamazsanız veya Windows kullanıyorsanız, birkaç seçeneğiniz var.

İlk seçeneğiniz, django'nun kaynak kodlarını indirip, `python setup.py install`
komutunu kaynak kodlarını açmış olduğunuz dizinden çalıştırarak kurulum yapmak.
Eğer windows kullanıcısı iseniz, python sistem yolunuzda olmayabilir. O yüzden
ya python'u sistem yolunuza ekleyin, ya da python yorumlayıcısının tam yolunu
kullanın.

Gnu/Linux sistemlerde ise bunu yapmak için yönetici haklarına ihtiyaç duyabilirsiniz.
Diğer bir seçeneğiniz ise, eğer setuptools python paketi sisteminizde yüklüyse,
easy_install django komutuyla yükleme yapmak. Bu komut django web çatısını sizin
için indirip kuracaktır. Diğer bir seçenek ise pip aracılığıyla yükleme yapmak.
Eğer pip python paketi sisteminizde yüklüyse, pip install django komutu aracılığıyla
da yükleme yapabilirsiniz. Django'nun kurulumunu yaptıktan sonra, python kabuğuna
girip, şu komutu vererek, django'nun dosylarının içe aktarılabildiğinden emin olun:
`import django`

Eğer bu komutu verdikten sonra hiçbir hata almazsanız, django ile web geliştirme
yapmaya hazırsınız demektir.

Django'nun kurulumunu yaptıktan sonra, eğer daha önceden başlanmış bir projeniz
yoksa, `django-admin.py startproject proje_adi` komutuyla yeni bir django projesine
başlayabilirsiniz. Bu komut size *proje_adi* isimli bir dizin içinde, 4 adet dosya oluşturacak.

 - **__init__.py** Boş bir dosyadır. Bu dosyayla neredeyse hiç bir işiniz
      olmayacak. Bu dizinin bir python paketi olduğu belirtmek için oradadır.
 - **manage.py** django-admin.py ile neredeyse aynı işi yapar, ancak projenizi
     python yoluna eklemek ve `DJANGO_SETTINGS_MODULE` çevre değişkenini
     ayarlamak gibi birkaç ek fonksiyonu vardır. Bu yüzden projenizin yönetimini
	 bu modül aracılığıyla yapacaksınız. Ne yaptığınızdan çok emin değilseniz,
	 bu dosyayı olduğu gibi bırakın.
 - **urls.py** Sitemizde hangi url'in nasıl sunulacağına ilişkin bilgi içerir.
    Daha detaylı bilgiye ilerleyen zamanlarda değineceğiz.
 - **settings.py** Projenin bütün ayarları bu modülün içerisindedir. manage.py
    ile aynı dizin içinde olması ve adının settings.py olması şarttır. Aksi halde
	başınız bir hayli ağrıyacaktır.Ayarların detaylarına birazdan değineceğiz.
   
Böylece içi boş bir django projesine başlamış olduk. Eğer gözlerinizle şahit
olmak isterseniz, proje dizininizin içindeyken,

{% codeblock lang:bash %}
python manage.py runserver
{% endcodeblock %}

komutunu vererek geliştirme sunucusunu başlatabilirsiniz. Bu sunucuyu
geliştirme süreci boyunca kullanacaksınız, ama günlük kullanım web sunucusu
olarak tavsiye edilmiyor. Eğer sunucunuz başarıyla çalıştıysa buraya tıklayarak
django'nun *It Worked* (Çalıştı) sayfasını görebilirsiniz. Ama henüz birşey yapmış
değiliz.

Bu yazının son kısmında biraz ayarlara bakacağız. Django projenizin ayarları
*settings.py* içerisinde bulunur. Bu modülü istediğiniz zaman uygulamalarınıza
`import` ile dahil ederek kullanabilirsiniz. Şimdi her ayara tek tek bakmaktansa
birkaç tanesinin üzerinde duracağım.

{% codeblock lang:python %}
DEBUG = True
{% endcodeblock %}

Bu ayar django projenizde bir hata olduğu zaman ayrıntılı hata ayıklama mesajlarını
görmenizi sağlıyor. Geliştirme sunucuzdayken çok kullanışlı bir özellik olsa da, günlük
kullanım sunucunuza geçtiğinizde kapatmanız gerekir. Günlük kullanım sunucunuzda açık
bırakmak güvenlik açığına neden olur.

{% codeblock lang:python %}
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.sqlite3', # Add 'postgresql_psycopg2', 'postgresql', 'mysql', 'sqlite3' or 'oracle'.
        'NAME': '/home/yasar/database.db',                      # Or path to database file if using sqlite3.
        'USER': '',                      # Not used with sqlite3.
        'PASSWORD': '',                  # Not used with sqlite3.
        'HOST': '',                      # Set to empty string for localhost. Not used with sqlite3.
        'PORT': '',                      # Set to empty string for default. Not used with sqlite3.
    }
}
{% endcodeblock %}

Burada veritabanı ayarlarını yapıyoruz. Blog uygulaması örneğimizde kullanımı
kolay olması açısından `sqlite3` veritabanı kullanacağız. Bu yüzden `ENGINE`'i
`django.db.backends.sqlite3` olarak ayarladık. Diğer seçeneklerimiz `sqlite3`
yerine `postgresql_psycopg2`, `postgresql`, `mysql` veya `oracle` olabilirdi.
Bunların hangisinin ne anlama geldiği yeteri kadar açık sanırım. `NAME` anahtarı
kullandığımız veritabanının adı, eğer `sqlite3` kullanıyorsak, bu veritabanımızın
dosya yolu. Tam yol kullanmayı unutmayın. `USER` ve `PASSWORD`, veritabanımızın
kullanıcı adı ve şifresi, `sqlite` için boş bırakabiliriz. `HOST` ve `PORT` ise
sırasıyla veritabanın hangi makinede bulunduğunu ve portunu belirtiyor. Eğer
veritabanı localhost'daysa ve öntanımlı port üzerindeyse, boş bırakabiliriz.
Ayrıca sqlite için bu değerlerin bir anlamı yok.

{% codeblock lang:python %}
TIME_ZONE = 'Europe/Istanbul'
{% endcodeblock %}

zaman dilimimizi ayarlıyoruz. Windows sistemlerde, sisteminizde ayarlı saat
dilimiyle aynı olmalı. Türkiye'deki windows kullanıcıları bunun için bu değeri
`Europe/Istanbul` olarak ayarlamaları gerekiyor. Unix sistemlerde ise None girmek
django'nun otomatik olarak sistem saatinizi kullanmasına neden olur. Buraya
girilebilecek değerlerin büyük bir çoğunluğunu burada bulabilirsiniz.

{% codeblock lang:python %}
LANGUAGE_CODE = 'tr'
{% endcodeblock %}

Bu projenin dilini ayarlıyor. Türkçe için `tr` girebilirsiniz.

{% codeblock lang:python %}
STATIC_ROOT = '/home/yasar/static_files'
{% endcodeblock %}

css, js ve resim dosyaları gibi içeriği sabit dosyaların toplanacağı yer. Buraya
kendiniz birşey eklemeyin. Detaylarına daha sonra değineceğiz. Bu dizini proje
dosyanızın dışarısında tutmakta fayda var.

{% codeblock lang:python %}
STATIC_URL = '/static/'
{% endcodeblock %}

Bu statik dosyaları hangi link üzerinden sunulacağını ayarlıyor. Olduğu gibi bırakılabilir.
{% codeblock lang:python %}
SECRET_KEY = 'asdfasdşfj qğwfasdfj lşahwe fhaşslkdnfj pqwdnf ş'
{% endcodeblock %}

Bunu kimseye göstermeyin. Yeteri kadar uzun, eşsiz ve rastgele bir şey olmalıdır.
Size özel bir gizli anahtardır. Şifreleme işlemleri için kullanılır. Ayrıca,
hazırda çalışan bir sitenin `SECRET_KEY`'ini değiştirmek, eski verilere ulaşamamanıza
neden olabilir. Projenizi ilk başlattığınızda size özel bir adet oluşturulur.
Olduğu gibi bırakabilirsiniz.

{% codeblock lang:python %}
INSTALLED_APPS = (
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.sites',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    # Uncomment the next line to enable the admin:
    # 'django.contrib.admin',
    # Uncomment the next line to enable admin documentation:
    # 'django.contrib.admindocs',
)
{% endcodeblock %}

Bu projede yüklü olan uygulamaların bir listesi. Bunlar yeni bir projeye
başladığınızda otomatik olarak eklenenler. Eğer bunlardan birkaçına ihtiyacınız
olmadığından eminseniz, buradan kaldırabilirsiniz. Ancak bunlara dokunmamayı
tavsiye ederim. En azından şimdilik. Bu yazıda inceleyeceğimiz ayarlar bu kadar.
Yazı dizisinin ilerleyen bölümlerinde yeri geldikçe ayarlar dosyasına tekrar döneceğiz