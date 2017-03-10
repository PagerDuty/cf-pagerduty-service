import os
import requests
import json
from flask import Flask

app = Flask(__name__)

port = os.getenv('PORT', '3000')
host = '0.0.0.0'


@app.route('/')
def index():
    return "Bind to pagerduty-service, then go to /trigger to trigger a PagerDuty incident"


@app.route('/trigger')
def trigger():
    post_trigger()
    return "Trigger incident was successful, page incoming..."


def get_env(env, default):
    try:
        v = os.environ[env]
    except:
        v = default
    return v


def config_path():
    default_path = "config.json"
    try:
        path = os.environ['DEMO_CONFIG_PATH']
    except:
        print "DEMO_CONFIG_PATH not set, using defaul_path"
        path = default_path
    return path


def parse_config(path):
    with open(path) as config_file:
        config = json.load(config_file)
    return config


def post_trigger():
    app_url = get_env('PAGERDUTY_API_URL', parse_config(config_path())['url'])
    username = get_env('PAGERDUTY_API_USERNAME', parse_config(config_path())['username'])
    password = get_env('PAGERDUTY_API_PASSWORD', parse_config(config_path())['password'])
    service_key = get_env('PAGERDUTY_API_TOKEN', parse_config(config_path())['token'])

    url = "http://{0}:{1}@{2}/pd/v1/trigger".format(username, password, app_url)
    headers = {
        "X-Broker-Api-Version": "2.11"
    }
    payload = {
        "service_key": service_key
    }

    r = requests.post(url, headers=headers, data=json.dumps(payload))
    r.raise_for_status()

if __name__ == "__main__":
    print "Listening on port {0}\n".format(port)
    app.run(host=host, port=port)
