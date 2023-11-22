package types

type (
	BaseEventData struct {
		Typ string `json:"typ"`
	}
	FocusEventData struct {
		BaseEventData
	}
	BlurEventData struct {
		BaseEventData
	}
	ScrollEventData struct {
		BaseEventData
	}
	KeyboardEventData struct {
		BaseEventData
		CtrlKey          bool   `json:"ctrlKey"`
		AltKey           bool   `json:"altKey"`
		ShiftKey         bool   `json:"shiftKey"`
		Key              string `json:"key"`
		Repeat           bool   `json:"repeat"`
		Code             string `json:"code"`
		IsComposing      bool   `json:"isComposing"`
		Location         int    `json:"location"`
		MetaKey          bool   `json:"metaKey"`
		Detail           int    `json:"detail"`
		EventTyp         string `json:"type"`
		TimeStamp        int    `json:"timeStamp"`
		Bubbles          bool   `json:"bubbles"`
		Cancelable       bool   `json:"cancelable"`
		Composed         bool   `json:"composed"`
		EventPhase       int    `json:"eventPhase"`
		IsTrusted        bool   `json:"isTrusted"`
		DefaultPrevented bool   `json:"defaultPrevented"`
	}
	MouseEventData struct {
		BaseEventData
		CtrlKey          bool   `json:"ctrlKey"`
		AltKey           bool   `json:"altKey"`
		ShiftKey         bool   `json:"shiftKey"`
		ClientX          int    `json:"clientX"`
		ClientY          int    `json:"clientY"`
		PageX            int    `json:"pageX"`
		PageY            int    `json:"pageY"`
		OffsetX          int    `json:"offsetX"`
		OffsetY          int    `json:"offsetY"`
		ScreenX          int    `json:"screenX"`
		ScreenY          int    `json:"screenY"`
		Buttons          int    `json:"buttons"`
		Button           int    `json:"button"`
		MovementX        int    `json:"movementX"`
		MovementY        int    `json:"movementY"`
		X                int    `json:"x"`
		Y                int    `json:"y"`
		Detail           int    `json:"detail"`
		EventTyp         string `json:"type"`
		TimeStamp        int    `json:"timeStamp"`
		Bubbles          bool   `json:"bubbles"`
		Cancelable       bool   `json:"cancelable"`
		Composed         bool   `json:"composed"`
		EventPhase       int    `json:"eventPhase"`
		IsTrusted        bool   `json:"isTrusted"`
		DefaultPrevented bool   `json:"defaultPrevented"`
	}
	ClickEventData struct {
		BaseEventData
		CtrlKey            bool    `json:"ctrlKey"`
		AltKey             bool    `json:"altKey"`
		ShiftKey           bool    `json:"shiftKey"`
		IsPrimary          bool    `json:"isPrimary"`
		ClientX            int     `json:"clientX"`
		ClientY            int     `json:"clientY"`
		Height             int     `json:"height"`
		Width              int     `json:"width"`
		PointerType        string  `json:"pointerType"`
		Pressure           float64 `json:"pressure"`
		TangentialPressure float64 `json:"tangentialPressure"`
		TiltX              int     `json:"tiltX"`
		TiltY              int     `json:"tiltY"`
		Twist              int     `json:"twist"`
	}
	OtherEventData struct {
		BaseEventData
		Name string `json:"name"`
	}
	EventData interface {
		ClickEventData | OtherEventData
	}
	Action interface {
		GetType() string
	}
	Renderer interface {
		Render() string
	}
	Component interface {
		Renderer
	}
	Page interface {
		Renderer
	}
)
