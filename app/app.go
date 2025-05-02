package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de Servidores na internet"

	// Adicionando os camondos
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			// Flag é um slice
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "biodoc.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	// Aqui usamos o pacote net para buscar os IPs
	// LookupIp vai retornar um slice porque dependendo do redirecionamento do site pode ser que ele tenha mais de um IP publico
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
