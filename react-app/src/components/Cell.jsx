import React from "react";

function Cell({ rIndex, cIndex, start, end, isPath, onClick }) {
  const getClassName = () => {
    if (start?.x === cIndex && start?.y === rIndex) return "cell start";
    if (end?.x === cIndex && end?.y === rIndex) return "cell end";
    if (isPath(rIndex, cIndex)) return "cell path";
    return "cell";
  };
  return (
    <div
      key={`${rIndex}-${cIndex}`}
      className={getClassName()}
      onClick={() => onClick(rIndex, cIndex)}
    ></div>
  );
}

export default Cell;
