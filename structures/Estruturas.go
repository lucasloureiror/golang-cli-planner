package structures

type Tarefas struct {
	nome string
}

type TarefasHoje struct {
	tarefas    []Tarefas
	quantidade int
	existe     bool
}

type TarefasProximas struct {
	tarefas    []Tarefas
	quantidade int
	pomodoro   int
	existe     bool
}

type TarefasDistantes struct {
	tarefas    []Tarefas
	quantidade int
	pomodoro   int
	existe     bool
}
type EstudoHoje struct {
	tarefas    []Tarefas
	quantidade int
	pomodoro   int
	existe     bool
}

type TrabalhoOuProva struct {
	tarefas    []Tarefas
	quantidade int
	pomodoro   int
	existe     bool
}
