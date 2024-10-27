import http.client
import json
from typing import List, Dict, Tuple, Any

def get_coworking_name() -> str:
    return input("Enter Coworking Name: ")

def get_number_of_workspaces() -> int:
    return int(input("How many workspaces do you want to add? "))

def get_workspace_data(workspace_index: int) -> Dict[str, Any]:
    workspace_name = input(f"Enter Workspace {workspace_index + 1} Name: ")
    workspace_seats = int(input(f"Enter Workspace {workspace_index + 1} Seats: "))
    return {"Name": workspace_name, "Seats": workspace_seats}

def create_coworking_payload(coworking_name: str, workspaces: List[Dict[str, Any]]) -> Dict[str, Any]:
    return {
        "Name": coworking_name,
        "Workspaces": workspaces
    }

def send_post_request(host: str, port: int, endpoint: str, json_data: str) -> Tuple[int, str]:
    conn = http.client.HTTPConnection(host, port)
    headers = {
        "Content-Type": "application/json",
        "Content-Length": str(len(json_data))
    }
    conn.request("POST", endpoint, json_data, headers)
    response = conn.getresponse()
    response_data = response.read()
    conn.close()
    return response.status, response_data.decode()

def main() -> None:
    coworking_name = get_coworking_name()
    num_workspaces = get_number_of_workspaces()
    workspaces: List[Dict[str, Any]] = []

    for i in range(num_workspaces):
        workspace = get_workspace_data(i)
        workspaces.append(workspace)

    payload = create_coworking_payload(coworking_name, workspaces)
    json_data = json.dumps(payload)
    print("Sending JSON data:", json_data)

    host = "localhost"
    port = 8080
    endpoint = "/api/coworking"

    status, response_data = send_post_request(host, port, endpoint, json_data)

    if status == 200:
        print("Coworking created successfully:", response_data)
    else:
        print("Error sending request:", status, response_data)

if __name__ == "__main__":
    main()
