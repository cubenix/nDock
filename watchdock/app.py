from flask import Flask
from config_manager import ConfigManager
import services.renderer as renderer


app = Flask(__name__)
hosts = ConfigManager().get_hosts()


@app.route('/')
@app.route('/index')
@app.route('/home')
@app.route('/index.html')
def index():
    return renderer.render_dashboard(hosts)


@app.route('/stats')
def stats():
    return renderer.render_stats(hosts)


if __name__ == '__main__':
    app.run(debug=True)
