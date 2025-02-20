import React, { useState } from "react";
import Cell from "./Cell";

const GridSize = 20;

function Grid() {
  const [grid, setGrid] = useState(
    Array(GridSize)
      .fill()
      .map(() => Array(GridSize).fill("1"))
  );
  const [start, setStart] = useState(null);
  const [end, setEnd] = useState(null);
  const [path, setPath] = useState([]);

  const [dfs, setDfs] = useState(true);

  const handleCellClick = (row, col) => {
    if (!start) {
      setStart({ x: col, y: row });
    } else if (!end) {
      const endPoint = { x: col, y: row };
      setEnd(endPoint);
      fetchPathRequest(start, endPoint);
    } else {
      setStart({ x: col, y: row });
      setEnd(null);
      setPath([]);
    }
  };

  const fetchPathRequest = async (start, end) => {
    if (!start || !end) return;

    try {
      const url = dfs
        ? "http://localhost:4000/find-path"
        : "http://localhost:4000/find-path-BFS";

      const res = await fetch(url, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ start: start, end: end }),
      });

      const data = await res.json();

      setPath(data.Path);
    } catch (error) {
      console.log({ e: error });
    }
  };

  const isPath = (row, col) => {
    return path?.some((p) => p.x === col && p.y === row) || false;
  };
  return (
    <div>
      <div>
        <button onClick={() => setDfs(!dfs)}>
          Switch {dfs ? "BFS" : "DFS"}
        </button>
        <div>Current algo</div>
        <p>{dfs ? "DFS" : "BFS"} </p>
      </div>

      <div className="grid-container">
        {grid.map((row, rIndex) =>
          row.map((cell, cIndex) => (
            <Cell
              key={`${rIndex}-${cIndex}`}
              rIndex={rIndex}
              cIndex={cIndex}
              start={start}
              end={end}
              isPath={isPath}
              onClick={handleCellClick}
            />
          ))
        )}
      </div>
    </div>
  );
}

export default Grid;
