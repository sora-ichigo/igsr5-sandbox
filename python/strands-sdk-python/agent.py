from strands import Agent
# https://github.com/strands-agents/tools/blob/8341c1126b36f61fc33c0ac4e0ccabe6495a3119/src/strands_tools/calculator.py#L590
from strands_tools import calculator
from strands.models.bedrock import BedrockModel
import hello_world

bedrock_model = BedrockModel(
    model_id="arn:aws:bedrock:ap-northeast-1:247679429177:inference-profile/apac.anthropic.claude-sonnet-4-20250514-v1:0",
    region_name='ap-northeast-1',
)

# TOOL_SPEC: https://github.com/strands-agents/tools/blob/8341c1126b36f61fc33c0ac4e0ccabe6495a3119/src/strands_tools/file_read.py#L234
# TOOL_SPEC 呼び出し元: https://github.com/strands-agents/sdk-python/blob/ffc7c5e68a50e0b8a9afb4c997a0eec8f88e4bf4/src/strands/tools/registry.py#L339
agent = Agent(tools=[calculator, hello_world], model=bedrock_model)
print(agent.model.config)

agent("What is the square root of 1764")

print("\n" + "="*50 + "\n")

# Test the hello_world tool
agent("Please use the hello world tool")

print("\n" + "="*50 + "\n")

# Direct method call to hello_world tool
result = agent.tool.hello_world()
print(f"Direct call result: {result}")
