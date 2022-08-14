package pb

import (
	"net"

	"gitee.com/healthy/backend/pkg/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCClient interface {
	Host() string
	Port() string
	User() UserManagerClient
	Personal() PersonalManagerClient
	CaseData() CaseDataManagerClient
	Class() ClassManagerClient
	MeasureDetail() MeasureDetailManagerClient
	MeasureResult() MeasureResultManagerClient
	MoveData() MoveDataManagerClient
	MovePrescription() MovePrescriptionManagerClient
	Work() WorkManagerClient
	WorkSubmit() WorkSubmitManagerClient
	Close() error
}

type grpcClient struct {
	host             string
	port             string
	conn             *grpc.ClientConn
	user             UserManagerClient
	personal         PersonalManagerClient
	caseData         CaseDataManagerClient
	class            ClassManagerClient
	measureDetail    MeasureDetailManagerClient
	measureResult    MeasureResultManagerClient
	moveData         MoveDataManagerClient
	movePrescription MovePrescriptionManagerClient
	work             WorkManagerClient
	workSubmit       WorkSubmitManagerClient
}

func NewGRPCClient() (GRPCClient, error) {
	host := config.GetString(config.ManagerHost)
	port := config.GetString(config.ManagerPort)
	address := net.JoinHostPort(host, port)

	logrus.Infoln("address", address)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		logrus.Errorf("did not connect: %v", err)
		return nil, err
	}

	return &grpcClient{
		host:             host,
		port:             port,
		conn:             conn,
		user:             NewUserManagerClient(conn),
		personal:         NewPersonalManagerClient(conn),
		caseData:         NewCaseDataManagerClient(conn),
		class:            NewClassManagerClient(conn),
		measureDetail:    NewMeasureDetailManagerClient(conn),
		measureResult:    NewMeasureResultManagerClient(conn),
		moveData:         NewMoveDataManagerClient(conn),
		movePrescription: NewMovePrescriptionManagerClient(conn),
		work:             NewWorkManagerClient(conn),
		workSubmit:       NewWorkSubmitManagerClient(conn),
	}, nil
}

func (c *grpcClient) Host() string {
	return c.host
}

func (c *grpcClient) Port() string {
	return c.port
}

func (c *grpcClient) User() UserManagerClient {
	return c.user
}

func (c *grpcClient) Personal() PersonalManagerClient {
	return c.personal
}

func (c *grpcClient) CaseData() CaseDataManagerClient {
	return c.caseData
}

func (c *grpcClient) Class() ClassManagerClient {
	return c.class
}

func (c *grpcClient) MeasureDetail() MeasureDetailManagerClient {
	return c.measureDetail
}

func (c *grpcClient) MeasureResult() MeasureResultManagerClient {
	return c.measureResult
}

func (c *grpcClient) MoveData() MoveDataManagerClient {
	return c.moveData
}

func (c *grpcClient) MovePrescription() MovePrescriptionManagerClient {
	return c.movePrescription
}

func (c *grpcClient) Work() WorkManagerClient {
	return c.work
}

func (c *grpcClient) WorkSubmit() WorkSubmitManagerClient {
	return c.workSubmit
}

func (c *grpcClient) Close() error {
	if c.conn != nil {
		logrus.Debugln("client: ", c.conn.GetState(), "->closed")
		return c.conn.Close()
	}
	return nil
}
