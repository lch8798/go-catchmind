import React, { Component } from "react";
import ReactDOM from "react-dom";
import io from "socket.io-client";

import "./index.css";

let picture = [];
const fps = 30;

const PORT = 3000;
const WEB_SOCKET_HOST = `ws://localhost:${PORT}`;

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      socket: io(WEB_SOCKET_HOST, { transports: ["websocket", "polling"] }),
      canvas: null,
      ctx: null,
      drawing: false,
    };

    // 실시간 드로잉
    this.state.socket.on("draw", async (data) => {
      picture.push(data);
    });

    // 방문 랜더링
    this.state.socket.on("drawInit", async (data) => {
      picture = data;
      this.rendering();
    });

    // 캔버스 초기화
    this.state.socket.on("paintInit", (data) => {
      const { canvas, ctx } = this.state;
      picture = data;
      ctx.fillStyle = "white";
      ctx.fillRect(0, 0, canvas.width, canvas.height);
    });
  }

  componentDidMount() {
    this.init();
  }

  init = async () => {
    const canvas = this.refs.canvas;
    const ctx = canvas.getContext("2d");
    ctx.fillStyle = "white";
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    await this.setState({ canvas });
    await this.setState({ ctx });

    const frame = fps / 1000;
    this.state.socket.emit("drawInit");
    setInterval(() => this.rendering(), frame);
  };

  paintInit = () => {
    this.state.socket.emit("paintInit");
  };

  rendering = () => {
    const { ctx } = this.state;
    ctx.fillStyle = "black";
    picture.forEach((location) => {
      ctx.fillRect(location[0], location[1], 2, 2);
    });
  };

  draw = (e) => {
    const { canvas, ctx, drawing, socket } = this.state;

    if (!drawing) {
      return;
    }

    let drawingLocation = [0, 0];

    if (e.touches) {
      // 터치
      drawingLocation = [
        e.touches[0].clientX - canvas.offsetLeft,
        e.touches[0].clientY - canvas.offsetTop,
      ];
    } else {
      // 마우스
      drawingLocation = [
        e.clientX - canvas.offsetLeft,
        e.clientY - canvas.offsetTop,
      ];
    }

    // 클라이언트 랜더링
    // picture.push(drawingLocation);

    socket.emit("draw", JSON.stringify(drawingLocation));
  };

  onDraw = () => {
    this.setState({ drawing: true });
    console.log("On!");
  };

  offDraw = () => {
    this.setState({ drawing: false });
    console.log("Off!");
  };

  render() {
    return (
      <div className="App">
        <button onClick={this.paintInit}>초기화</button>
        <div className="game" id="game">
          <canvas
            id="canvas"
            ref="canvas"
            width={600}
            height={600}
            onMouseUp={this.offDraw}
            onMouseDown={this.onDraw}
            onMouseMove={this.draw}
            onTouchStart={this.onDraw}
            onTouchEnd={this.offDraw}
            onTouchMove={this.draw}
          />
        </div>
      </div>
    );
  }
}

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById("root")
);
