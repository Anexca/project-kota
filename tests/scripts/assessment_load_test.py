import requests
import time
from concurrent.futures import ThreadPoolExecutor

server_url = "https://server-dev-1082143187273.asia-south1.run.app"

urls = [
    f"{server_url}/exams/banking/descriptive/252/evaluate?isopen=false",
    f"{server_url}/exams/banking/descriptive/251/evaluate?isopen=false",
    f"{server_url}/exams/banking/descriptive/187/evaluate?isopen=false"
]

bearer_token = "##SECRET##"

headers = {
    "Authorization": f"Bearer {bearer_token}",
    "Content-Type": "application/json"
}

data = {
    "completed_seconds": 0,
    "content": "In an increasingly interconnected yet volatile global market, export credit agencies (ECAs) play a crucial role in facilitating and promoting international trade. These institutions provide a safety net for exporters by offering insurance against a range of political and commercial risks, including non-payment by foreign buyers, political instability in buyer countries, and currency fluctuations..."
}

def poll_assessment(id):
    assessment_url = f"{server_url}/exams/assesments/{id}"
    
    while True:
        response = requests.get(assessment_url, headers=headers)
        if response.status_code == 200:
            assessment_data = response.json()
            status = assessment_data['data']['status']
            print(f"Assessment {id} status: {status}")
            if status != "PENDING":
                print(f"Assessment {id} completed with status: {status}")
                break
        else:
            print(f"Failed to retrieve assessment {id}: {response.text}")
            break
        time.sleep(5)  # Poll every 5 seconds

def call_api(url, attempt):
    print(f"Calling {url} (Attempt {attempt})")
    response = requests.post(url, headers=headers, json=data)
    
    if response.status_code == 200 or response.status_code == 202:
        result = response.json()
        assessment_id = result['data']['id']
        print(f"Initial request successful for {url} (Attempt {attempt}). Assessment ID: {assessment_id}")
        # Start polling the assessment endpoint
        poll_assessment(assessment_id)
    else:
        print(f"Failed with status code {response.status_code}: {response.text}")

with ThreadPoolExecutor(max_workers=6) as executor:
    futures = []
    for url in urls:
        for i in range(2):  # Call the API twice for each URL
            futures.append(executor.submit(call_api, url, i + 1))

    for future in futures:
        future.result()  # Wait for the result (blocks until all tasks are done)