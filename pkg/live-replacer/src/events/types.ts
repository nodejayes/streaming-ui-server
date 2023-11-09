export interface BaseEventData {
    typ: string;
}

export interface ClickEventData extends BaseEventData{
    ctrlKey: boolean;
    altKey: boolean;
    shiftKey: boolean;
    isPrimary: boolean;
    clientX: number;
    clientY: number;
    height: number;
    width: number;
    pointerType: string;
    pressure: number;
    tangentialPressure: number;
    tiltX: number;
    tiltY: number;
    twist: number;
    button: number;
    buttons: number;
    metaKey: boolean;
    movementX: number;
    movementY: number;
    offsetX: number;
    offsetY: number;
    pageX: number;
    pageY: number;
    screenX: number;
    screenY: number;
    x: number;
    y: number;
}

export interface MouseEventData extends BaseEventData {
    ctrlKey: boolean;
    altKey: boolean;
    shiftKey: boolean;
    clientX: number;
    clientY: number;
    pageX: number;
    pageY: number;
    offsetX: number;
    offsetY: number;
    screenX: number;
    screenY: number;
    buttons: number;
    button: number;
    movementX: number;
    movementY: number;
    x: number;
    y: number;
    detail: number;
    type: string;
    timeStamp: number;
    bubbles: boolean;
    cancelable: boolean;
    composed: boolean;
    eventPhase: number;
    isTrusted: boolean;
    defaultPrevented: boolean;
}

export interface KeyboardEventData extends BaseEventData {
    ctrlKey: boolean;
    altKey: boolean;
    shiftKey: boolean;
    key: string;
    repeat: boolean;
    code: string;
    isComposing: boolean;
    location: number;
    metaKey: boolean;
    detail: number;
    type: string;
    timeStamp: number;
    bubbles: boolean;
    cancelable: boolean;
    composed: boolean;
    eventPhase: number;
    isTrusted: boolean;
    defaultPrevented: boolean;
}

export type EventData = ClickEventData | MouseEventData | KeyboardEventData;