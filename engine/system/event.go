package system

import (
	eventpb "github.com/Intrising/intri-type/event"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func sendBootEvent(action eventpb.BootActionTypeOptions, version string) {
	var msg string
	var loggingType eventpb.LoggingTypeOptions
	if action == eventpb.BootActionTypeOptions_BOOT_ACTION_TYPE_WARM_START {
		msg = "System Warmstart. Reason: Reboot. Version Information: " + version
		loggingType = eventpb.LoggingTypeOptions_LOGGING_TYPE_BOOT_WARM_START
	} else if action == eventpb.BootActionTypeOptions_BOOT_ACTION_TYPE_COLD_START {
		msg = "Initial System Cold Start. Reason: Power Up. Version Information: " + version
		loggingType = eventpb.LoggingTypeOptions_LOGGING_TYPE_BOOT_COLD_START
	} else {
		msg = "all protocols are ready"
		loggingType = eventpb.LoggingTypeOptions_LOGGING_TYPE_NONE
	}

	// log.Println("sendBootEvent : ", msg, loggingType)

	evt := &eventpb.Internal{
		Ts:          timestamppb.Now(),
		Type:        eventpb.InternalTypeOptions_INTERNAL_TYPE_BOOT,
		Message:     msg,
		LoggingType: loggingType,
		Parameter: &eventpb.Internal_Boot{
			Boot: &eventpb.BootParameter{
				Type:    action,
				Version: version,
			},
		},
	}
	halClient.EventClient.SendEvent(evt)
}

func sendResetButtonEvent(in eventpb.ButtonTriggerActionTypeOptions) {
	var msg string
	if in == eventpb.ButtonTriggerActionTypeOptions_BUTTON_TRIGGER_ACTION_TYPE_REBOOT {
		msg = "Reboot the device by press the system button"
	} else { // in == eventpb.ButtonTriggerActionTypeOptions_BUTTON_TRIGGER_ACTION_TYPE_FACTORY
		msg = "Factory Default by press the system button"
	}

	evt := &eventpb.Internal{
		Ts:          timestamppb.Now(),
		Type:        eventpb.InternalTypeOptions_INTERNAL_TYPE_BUTTON,
		Message:     msg,
		LoggingType: eventpb.LoggingTypeOptions_LOGGING_TYPE_BUTTON_PRESSED,
		Parameter: &eventpb.Internal_Button{
			Button: &eventpb.ButtonParameter{
				Type:    eventpb.ButtonTypeOptions_BUTTON_TYPE_RESET,
				Action:  eventpb.ButtonActionTypeOptions_BUTTON_ACTION_TYPE_PRESSED,
				Trigger: in,
			},
		},
	}

	halClient.EventClient.SendEvent(evt)
}
