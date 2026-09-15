// Q: 5
const getChaseVerdict = (target, scored, ballsLeft) => {
  let runsNeeded = target - scored;
  let result = "";
  if (runsNeeded <= 0) {
    return "Won";
  } else if (ballsLeft <= 0) {
    return "Lost";
  } else {
    let requiredRate = (runsNeeded / ballsLeft) * 6;
    if (requiredRate <= 6) {
      result = "Comfortable";
    } else if (requiredRate > 6 && requiredRate <= 12) {
      result = "Tough";
    } else {
      result = "Almost Impossible";
    }
  }
  return `"Need ${runsNeeded} run in ${ballsLeft} balls | ${result}"`;
};

console.log(getChaseVerdict(150, 149, 1));
