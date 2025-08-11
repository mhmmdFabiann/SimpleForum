package configs

import "github.com/spf13/viper"

// Pointer global ke struct Config hasil unmarshal dari file config
var config *Config

// Struct internal untuk menyimpan opsi konfigurasi
type option struct {
	configFolders []string // Folder tempat mencari file konfigurasi
	configFile    string   // Nama file konfigurasi (tanpa ekstensi)
	configType    string   // Tipe file (yaml, json, dll)
}

// Init adalah fungsi utama untuk menginisialisasi konfigurasi
// Ia menerima variadic parameter Option (fungsi-fungsi yang mengubah struct option)
func Init(opts ...Option) error {
	// Set nilai default
	opt := &option{
		configFolders: getDefaultConfigFolders(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	// Jalankan fungsi-fungsi opsional yang bisa overwrite default
	for _, optFunc := range opts {
		optFunc(opt)
		// optFunc adalah fungsi yang memodifikasi nilai-nilai dalam struct option
	}

	// Tambahkan semua folder yang bisa berisi file konfigurasi
	for _, confFolder := range opt.configFolders {
		viper.AddConfigPath(confFolder)
	}

	// Set nama file dan tipe file ke Viper
	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)

	// Baca environment variable secara otomatis
	viper.AutomaticEnv()

	// Buat instance kosong dari Config (struct ini harus kamu definisikan di file lain)
	config = new(Config)

	// Baca file konfigurasi dari path yang telah ditentukan
	err := viper.ReadInConfig()
	if err != nil {
		return err // return error kalau file tidak bisa dibaca
	}

	// Mapping isi file ke struct Config
	return viper.Unmarshal(config)
}

// Option adalah tipe fungsi yang mengubah pointer ke struct option
// Misalnya: func(o *option) { o.configFile = "dev_config" }
type Option func(*option)

// Fungsi-fungsi default value:
func getDefaultConfigFolders() []string {
	return []string{"./configs"} // default folder
}

func getDefaultConfigFile() string {
	return "config" // default filename: config.yaml
}

func getDefaultConfigType() string {
	return "yaml" // default filetype
}

// func untuk melakukan overwrite 
func WithConfFolder(configFolders []string) Option{
	return func(o *option) {
		o.configFolders = configFolders
	}
}

func WithConfFile(configFile string) Option{
	return func(o *option) {
		o.configFile = configFile
	}
}

func WithConfType(configType string) Option{
	return func(o *option) {
		o.configType = configType
	}
}

func GetConf() *Config{
	if config == nil{
		config = &Config{}
	}
	return config
}