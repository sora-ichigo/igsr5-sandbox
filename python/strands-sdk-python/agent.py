from strands import Agent
# https://github.com/strands-agents/tools/blob/8341c1126b36f61fc33c0ac4e0ccabe6495a3119/src/strands_tools/calculator.py#L590
from strands_tools import calculator
from strands.models.bedrock import BedrockModel

bedrock_model = BedrockModel(
    model_id="arn:aws:bedrock:ap-northeast-1:247679429177:inference-profile/apac.anthropic.claude-sonnet-4-20250514-v1:0",
    region_name='ap-northeast-1',
)

agent = Agent(tools=[calculator], model=bedrock_model)
print(agent.model.config)

agent("What is the square root of 1764")
