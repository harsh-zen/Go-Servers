package handler

import (
	pb "native_server/protos"
	"net/http"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetAppointmentsProto( w http.ResponseWriter, r *http.Request){

	var protoAppointmentList pb.AppointmentList

	for _, appointment := range appointments {
        protoAppointment := &pb.Appointment{
            Id:        int32(appointment.ID),
            Title:     appointment.Title,
            StartTime: timestamppb.New(appointment.StartTime),
            EndTime:   timestamppb.New(appointment.EndTime),
        }
        protoAppointmentList.Appointments = append(protoAppointmentList.Appointments, protoAppointment)
    }
	resp, err := proto.Marshal(&protoAppointmentList)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(resp)
}

