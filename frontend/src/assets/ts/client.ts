import * as F from "../../../../backend/public/gomarvin.gen"

/** Pass in the auth token to set it in the header */
export function clientWithAuth(authToken: string): F.Client {
    const init = F.defaultClient

    const client: F.Client = {
        api_prefix: init.api_prefix,
        host_url: init.host_url,
        headers: {
            "Content-type": "application/json;charset=UTF-8",
            "Authorization": `Bearer ${authToken}`
        }
    }

    return client
}