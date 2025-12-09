package cmd

import (
	"log/slog"
	"os"

	"github.com/rodrigoconcham/gocodecli/mod2/healtcheck/logger"
	"github.com/spf13/cobra"
)

var (
	logFile string
	l       *slog.Logger
	//	threshold float64
	//	retries   int
	silent bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Una herramienta para verificar el estado de salud y capacidad de respuesta de aplicaciones web",
	Long: `El comando healthcheck está diseñado para evaluar la salud y 
capacidad de respuesta de aplicaciones web especificadas. Envía peticiones HTTP 
a las URLs proporcionadas por el usuario, evaluando si los servicios son 
accesibles y qué tan rápido responden. Este comando soporta tanto 
verificaciones inmediatas de una sola vez como monitoreo continuo, permitiendo a los usuarios 
especificar intervalos para evaluaciones de salud en curso. Con banderas (flags) 
adicionales para personalización, los usuarios pueden adaptar el comando para satisfacer diversas 
necesidades de monitoreo, desde simples verificaciones de disponibilidad hasta análisis detallados 
de rendimiento.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		l = logger.NewLogger(logFile)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logFile, "logfile", "healthcheck.log", "File to log output to")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().Float64Var(&threshold, "threshold", 0.5, "Threshold value for considering a response")
	rootCmd.PersistentFlags().IntVar(&retries, "retries", 3, "Number of retries for a failed request")
	rootCmd.PersistentFlags().BoolVar(&silent, "silent", false, "Run in silent mode without stdout output")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Run in verbose mode. Overrides silent mode")

}
