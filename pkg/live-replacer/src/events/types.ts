export interface BaseEventData {
    typ: string;
}

export interface ClickEventData extends BaseEventData{
    ctrlKey: boolean;
}

export type EventData = ClickEventData;