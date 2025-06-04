from strands import Agent
from strands_tools import calculator
from strands.models.bedrock import BedrockModel

bedrock_model = BedrockModel(
    model_id="arn:aws:bedrock:ap-northeast-1:247679429177:inference-profile/apac.anthropic.claude-sonnet-4-20250514-v1:0",
    region_name='ap-northeast-1',
)

agent = Agent(tools=[calculator], model=bedrock_model)
print(agent.model.config)

agent("What is the square root of 1764")
