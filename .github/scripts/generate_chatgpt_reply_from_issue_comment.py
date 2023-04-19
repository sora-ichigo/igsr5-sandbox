import os
import json
import openai
from github import Github

# Authenticate with the OpenAI API
openai.api_key = os.environ["OPENAI_API_KEY"]

# Authenticate with the GitHub API
g = Github(os.environ["GITHUB_TOKEN"])

# Get the issue comment event payload
event = os.environ["GITHUB_EVENT_PATH"]
with open(event, "r") as f:
    event_payload = json.load(f)

comment_body = event_payload["comment"]["body"]

# Check if @gptalk is mentioned in the comment
if "@gptalk" in comment_body:
    # Generate a response using the OpenAI API
    msg = f"{comment_body.replace('@gptalk', '')}"
    msgs = [
        {
        "role": "user",
        "content": msg
        }
    ]
    response = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=msgs
    )

    # Get the response text
    response_text = response.choices[0].message.content

    # Get the repository and issue from the event payload
    repo = g.get_repo(event_payload["repository"]["full_name"])
    issue = repo.get_issue(number=event_payload["issue"]["number"])

    # Comment on the issue with the generated response
    issue.create_comment(response_text)
