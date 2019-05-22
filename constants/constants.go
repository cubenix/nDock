package constants

// Host ports for web client and gRPC server
const (
	ClientPort = ":8000"
	ServerPort = ":5000"
)

// DockerAPIPort is the default TCP port on each host, at which Docker should be listening for TCP requests
const DockerAPIPort = ":2375"

// DockerAPIProtocol is the default protocol used to communicate with Docker hosts
const DockerAPIProtocol = "tcp://"

// DockerAPIVersion is the default API version used while creating a Docker client
const DockerAPIVersion = "v1.38"

// ContainerRunning is the running state of a container
const ContainerRunning = "running"

// BGClasses is a set background classes
var BGClasses = [6]string{
	bgPrimary,
	bgSuccess,
	bgInfo,
	bgWarning,
	bgDanger,
	bgSecondary,
}

// BGCodes is a set background color codes
var BGCodes = [6]string{
	bgColor4e73df,
	bgColor1cc88a,
	bgColor36b9cc,
	bgColorf6c23e,
	bgColore74a3b,
	bgColor858796,
}

// TextClasses is a set of text classes
var TextClasses = [6]string{
	textPrimary,
	textSuccess,
	textInfo,
	textWarning,
	textDanger,
	textSecondary,
}

const (
	bgPrimary   = "bg-primary"
	bgSuccess   = "bg-success"
	bgInfo      = "bg-info"
	bgWarning   = "bg-warning"
	bgDanger    = "bg-danger"
	bgSecondary = "bg-secondary"
)

const (
	bgColor4e73df = "#4e73df"
	bgColor1cc88a = "#1cc88a"
	bgColor36b9cc = "#36b9cc"
	bgColorf6c23e = "#f6c23e"
	bgColore74a3b = "#e74a3b"
	bgColor858796 = "#858796"
)

const (
	textPrimary   = "text-primary"
	textSuccess   = "text-success"
	textInfo      = "text-info"
	textWarning   = "text-warning"
	textDanger    = "text-danger"
	textSecondary = "text-secondary"
)
