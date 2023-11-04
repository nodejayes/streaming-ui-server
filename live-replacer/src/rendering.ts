import { ClickEvent } from "./events/click.event";

export async function render(api: WsApi) {
    new ClickEvent(api);
}