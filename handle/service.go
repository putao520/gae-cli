package handle

import (
	"fmt"
	"gae-cli/config"
	"gae-cli/gsc"
	"gae-cli/gsc/services"
	"github.com/eiannone/keyboard"
	"os"
	"strconv"
	"strings"
)

func waitSelectApp(appService *services.AppService) services.AppItem {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		appService.ClearPrint()
		if char == 'N' || char == 'n' {
			appService.Next()
			appService.Print()
			continue
		}
		if char == 'P' || char == 'p' {
			appService.Prev()
			appService.Print()
			continue
		}
		if char >= 48 && char <= 57 {
			println("selected " + string(char))
			return appService.Select(int(char - 48))
		}
	}
}

func selectApp() services.AppItem {
	// 显示当前账号在平台的所有应用
	appService := services.NewAppService()
	appService.Load()
	println("AppService:(please input service number to select service), press N to next page, press P to prev page, default is 0")
	appService.Print()
	// 如果键盘按下 n 按键，继续加载数据, 如果按下数字键，选择应用
	return waitSelectApp(appService)
}

func waitSelectService(serviceService *services.ServiceService) services.ServiceItem {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		serviceService.ClearPrint()
		if char == 'N' || char == 'n' {
			serviceService.Next()
			serviceService.Print()
			continue
		}
		if char == 'P' || char == 'p' {
			serviceService.Prev()
			serviceService.Print()
			continue
		}
		if char >= 48 && char <= 57 {
			println("selected " + string(char))
			return serviceService.Select(int(char - 48))
		}
	}
}

func selectService() services.ServiceItem {
	// 显示当前账号在平台的所有服务
	serviceService := services.NewServiceService()
	serviceService.Load()
	println("ServiceService:(please input service number to select service), press N to next page, press P to prev page")
	serviceService.Print()
	return waitSelectService(serviceService)
}

func waitDockerRegistry() string {
	println("please input docker registry, like is docker.io/username")
	print("docker registry url:")
	var input string
	fmt.Scanln(&input)
	return input
}

func MakeFolder(currentPath string, name string) {
	targetPath := currentPath + "/" + name
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		err := os.MkdirAll(targetPath, os.ModePerm)
		if err != nil {
			println("create " + name + " folder error")
			return
		}
	}
}

// 生成 main.java 文件
func MakeMainJava(currentPath string, serviceName string) {
	MakeFolder(currentPath, "/src/main/java/Main")
	file, err := os.Create(currentPath + "/src/main/java/Main/main.java")
	if err != nil {
		println("create Main.java error")
		return
	}
	defer file.Close()
	file.WriteString("package main.java.Main;\n\n")
	file.WriteString("import common.java.JGrapeSystem.GscBooster;\n\n")
	file.WriteString("public class Main {\n")
	file.WriteString("\tpublic static void main(String[] args) {\n")
	file.WriteString("\t\tSystem.out.println(\"<" + serviceName + "> service\");\n")
	file.WriteString("\t\tGscBooster.start(args, () -> {\n")
	file.WriteString("\t\t\t // load todo\n")
	file.WriteString("\t\t});\n")
	file.WriteString("\t}\n")
	file.WriteString("}\n")
}

func MakeDockerfile(currentPath string) {
	file, err := os.Create(currentPath + "/Dockerfile")
	if err != nil {
		println("create Dockerfile error")
		return
	}
	defer file.Close()
	file.WriteString("FROM ghcr.io/graalvm/graalvm-community:21\n")
	file.WriteString("COPY ./target/#{f} /home/app/\n")
	file.WriteString("WORKDIR /home/app\n")
	file.WriteString("ENTRYPOINT [\"java\", \"-Dfile.encoding=utf-8\", \"-jar\", \"#{f}\"]\n")
}

func MakeGaeShell(currentPath string) {
	// 创建 .gaeshell 文件
	file, err := os.Create(currentPath + "/.gaeshell")
	if err != nil {
		println("create .gaeshell error")
		return
	}
	defer file.Close()
	file.WriteString("echo \"构建jar\"\n")
	file.WriteString("mvn clean deploy \"-Dmaven.test.skip=true\"\n")
	file.WriteString("echo \"构建docker\"\n")
	file.WriteString("docker build --no-cache -t #{n}:#{v} .\n")
	file.WriteString("echo \"增加tags\"\n")
	file.WriteString("docker tag #{n}:#{v} #{h}#{n}:#{v}\n")
	file.WriteString("echo \"推送docker\"\n")
	file.WriteString("docker push #{h}#{n}:#{v}\n")
	file.WriteString("echo \"拉取docker命令\"\n")
	file.WriteString("echo \"docker pull #{h}#{n}:#{v}\"\n")
}

func MakeShell(currentPath string, version string, name string, host string) {
	// 判断是 linux 还是 windows
	shellName := "build.sh"
	if os.IsPathSeparator('\\') {
		shellName = "build.bat"
	}

	// 创建 shell 文件
	file, err := os.Create(currentPath + "/" + shellName)
	if err != nil {
		println("create build.sh error")
		return
	}
	defer file.Close()
	file.WriteString("gae run -v " + version + " -n " + name + " -h " + host + " -f #{n}-#{v}-jar-with-dependencies.jar\n")
}

func MakeRunnerShell(currentPath string, version string, name string, host string, token string, port int, appId string, deployName string) {
	// 判断是 linux 还是 windows
	shellName := "runner.sh"
	if os.IsPathSeparator('\\') {
		shellName = "runner.bat"
	}

	// 创建 shell 文件
	file, err := os.Create(currentPath + "/" + shellName)
	if err != nil {
		println("create runner.sh error")
		return
	}
	defer file.Close()
	file.WriteString("java -Dfile.encoding=utf-8 -jar " + name + "-" + version + "-jar-with-dependencies.jar -n " + deployName + " -h " + host + " -k " + token + " -p " + strconv.Itoa(port) + " -a " + appId + "\n")
}

func MakePomXML(currentPath string, serviceID string) {
	// 获得 pom.xml 文件内容
	res := gsc.GetRpcMasterClient().SetEndpoint("system", "Services").Call("downloadPom", []string{serviceID})
	if res == nil {
		println("download pom.xml error")
		return
	}
	pomXML := res.Body.Data.(string)
	file, err := os.Create(currentPath + "/pom.xml")
	if err != nil {
		println("create pom.xml error")
		return
	}
	defer file.Close()
	file.WriteString(pomXML)
}

func MakeGscProjectCfg(currentPath string, fullMasterUri string, app *services.AppItem, serv *services.ServiceItem, deployName string) {
	// 创建 .gaeshell 文件
	file, err := os.Create(currentPath + "/gfw.cfg")
	if err != nil {
		println("create gfw.cfg error")
		return
	}
	defer file.Close()
	// 通过 : 分割字符串
	masterPort := "80"
	masterUriArray := strings.Split(fullMasterUri, ":")
	masterHost := strings.Join(masterUriArray[:2], ":")
	if len(masterUriArray) >= 3 {
		masterPort = masterUriArray[2]
	}

	// 服务名 name
	file.WriteString("name=" + deployName + "\n")
	// 控制平面uri masterHost+masterPort
	file.WriteString("MasterHost=" + masterHost + "\n")
	file.WriteString("MasterPort=" + masterPort + "\n")
	// 控制平面token publicKey
	file.WriteString("publicKey=" + app.Secret + "\n")
	// 服务端口 port
	file.WriteString("port=" + strconv.Itoa(serv.Port) + "\n")
	// 应用id
	file.WriteString("appID=" + app.ID + "\n")
}

// 获得应用和服务信息
func ArgService() {
	println("init service")
	println("1: define service in development platform")
	println("2: deploy service in a application development platform")
	println("3: you can use gae init command to create a noobdog project in the current directory")
	println("----------------")

	provider := config.ReadGaeConfig().DefaultProvider
	appInfo := selectApp()
	println("selected app: " + appInfo.Name)
	serviceInfo := selectService()
	println("selected service: " + serviceInfo.Name)
	dockerRegistry := waitDockerRegistry()

	// 当前目录下创建一个 project 配置文件
	projectConfig := &config.GaeProjectConfig{
		Project: config.GaeProjectConfigInfo{
			Provider:       provider,
			AppID:          appInfo.ID,
			ServiceID:      strconv.Itoa(serviceInfo.ID),
			DockerRegistry: dockerRegistry,
		},
	}

	currentPath, err := os.Getwd()
	if err != nil {
		println("get current path error")
		return
	}

	deployInfo := services.QueryDeployInfo(appInfo.ID, strconv.Itoa(serviceInfo.ID))
	if deployInfo == nil {
		println("this service[" + serviceInfo.Name + "] is not deploy on application [" + appInfo.Name + "], please deploy first!")
		return
	}

	fullMasterUri := config.NewGaeConfigManager().GetDefaultProvider().Host

	println("create project files")
	config.CreateOrUpdateGaeProjectConfig(*projectConfig)
	println("folder building...")
	// 当前目录下创建 target 文件夹
	MakeFolder(currentPath, "target")
	// 当前目录下创建 test 文件夹
	MakeFolder(currentPath, "test")
	// 创建 main.java 文件
	println("main.java building...")
	MakeMainJava(currentPath, serviceInfo.Name)
	// 创建 dockerfile 文件
	println("dockerfile building...")
	MakeDockerfile(currentPath)
	// 创建 .gaeshell 文件
	println("shell script building...")
	MakeGaeShell(currentPath)
	MakeRunnerShell(currentPath, serviceInfo.Version, serviceInfo.Name, fullMasterUri, appInfo.Secret, serviceInfo.Port, appInfo.ID, deployInfo.Name)
	// 创建 shell 文件
	MakeShell(currentPath, serviceInfo.Version, serviceInfo.Name, dockerRegistry)
	// 创建 pom.xml 文件
	println("pom.xml building...")
	MakePomXML(currentPath, strconv.Itoa(serviceInfo.ID))
	// 创建 gfw.cfg 文件
	println("gfw.cfg building...")
	MakeGscProjectCfg(currentPath, fullMasterUri, &appInfo, &serviceInfo, deployInfo.Name)
	println("success!")
}
