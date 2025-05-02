package app

import "github.com/urfave/cli"

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
		},
	}

	return app
}
