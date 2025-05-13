# GO-Minus Vulkan Bağlayıcı Kütüphanesi

Bu paket, GO-Minus programlama dili için Vulkan API bağlayıcılarını içerir. Vulkan, modern grafik ve hesaplama API'sidir ve GO-Minus ile Vulkan kullanarak yüksek performanslı grafik uygulamaları geliştirilebilir.

## Özellikler

- Vulkan 1.3 API desteği
- GO-Minus'un nesne yönelimli programlama özelliklerini kullanarak temiz bir API
- Otomatik kaynak yönetimi
- Hata işleme ve istisna desteği
- Kapsamlı belgelendirme ve örnekler

## Gereksinimler

- GO-Minus 0.5.0 veya üzeri
- Vulkan SDK 1.3.0 veya üzeri
- Desteklenen bir GPU ve sürücü

## Kurulum

### Vulkan SDK Kurulumu

1. [Vulkan SDK](https://vulkan.lunarg.com/sdk/home) adresinden işletim sisteminize uygun SDK'yı indirin ve kurun.
2. Ortam değişkenlerinin doğru şekilde ayarlandığından emin olun.

### Paket Kullanımı

GO-Minus projenizde Vulkan paketini kullanmak için, `gom.mod` dosyanıza aşağıdaki bağımlılığı ekleyin:

```
require (
    vulkan v1.0.0
)
```

## Kullanım

```go
import (
    "vulkan"
)

func main() {
    // Vulkan instance oluşturma
    appInfo := vulkan.ApplicationInfo{
        sType: vulkan.STRUCTURE_TYPE_APPLICATION_INFO,
        pApplicationName: "GO-Minus Vulkan App",
        applicationVersion: vulkan.MAKE_VERSION(1, 0, 0),
        pEngineName: "No Engine",
        engineVersion: vulkan.MAKE_VERSION(1, 0, 0),
        apiVersion: vulkan.API_VERSION_1_0
    }
    
    createInfo := vulkan.InstanceCreateInfo{
        sType: vulkan.STRUCTURE_TYPE_INSTANCE_CREATE_INFO,
        pApplicationInfo: &appInfo
    }
    
    instance := vulkan.Instance{}
    result := vulkan.CreateInstance(&createInfo, null, &instance)
    
    if result != vulkan.SUCCESS {
        // Hata işleme
    }
    
    // Kullanım sonrası temizleme
    instance.destroy()
}
```

## Örnekler

- [Hello Triangle](../../examples/vulkan/hello-triangle/): Basit bir üçgen çizme örneği
- [Texture Loading](../../examples/vulkan/texture-loading/): Doku yükleme ve gösterme örneği
- [Compute Shader](../../examples/vulkan/compute-shader/): Hesaplama shader'ı kullanma örneği

## Belgelendirme

Daha fazla bilgi için [Vulkan API Referansı](../../docs/reference/vulkan.md) belgesine bakın.

## Lisans

GO-Minus Vulkan bağlayıcı kütüphanesi, [MIT Lisansı](../../LICENSE) altında lisanslanmıştır.
