import { ClickEvent } from "./events/click.event";

export async function render(api: WsApi) {
    ClickEvent.register(api);
}