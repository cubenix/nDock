from flask import Flask
from config_manager import ConfigManager
import services.renderer as renderer


app = Flask(__name__)
config_manager = ConfigManager()


@app.route('/')
@app.route('/index')
@app.route('/home')
@app.route('/index.html')
def index():
    return renderer.render_dashboard(config_manager)


if __name__ == '__main__':
    app.run(debug=True)
