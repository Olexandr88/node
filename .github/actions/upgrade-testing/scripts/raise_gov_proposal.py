import os
import requests
os.environ['NODE'] = "http://127.0.0.1:26657"
CURRENT_HEIGHT = requests.get(f"{os.environ['NODE']}/status").json()["result"]["sync_info"]["latest_block_height"]
UPGRADE_HEIGHT = int(CURRENT_HEIGHT) + (int(os.environ['PROPOSAL_TIME_SECONDS']) / int(os.environ['BLOCK_TIME_SECONDS']))
GOV_PROPOSAL = f"""zetacored tx gov submit-legacy-proposal software-upgrade "{os.environ['VERSION']}" \
    --from "{os.environ['MONIKER']}" \
    --deposit {os.environ["DEPOSIT"]} \
    --upgrade-height "{UPGRADE_HEIGHT}" \
    --upgrade-info "{os.environ['UPGRADE_INFO']}" \
    --title "{os.environ['VERSION']}" \
    --description "Zeta Release {os.environ['VERSION']}" \
    --chain-id "{os.environ['CHAINID']}" \
    --node "{os.environ['NODE']}" \
    --keyring-backend test \
    --fees {os.environ['FEES']} \
    -y \
    --no-validate"""
print(GOV_PROPOSAL)