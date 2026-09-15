// Q4: Dhaka CNG Fare Meter

function getCngFare(distance, isNight = false, waitingMinites = 0) {
  let totalFare = 0;
  let remainingDistance = 0;
  if (distance <= 2) {
    totalFare = 50 + waitingMinites * 2;
  } else if (distance > 2) {
    remainingDistance = distance - 2;
    totalFare = 50 + remainingDistance * 15 + waitingMinites * 2;
  }

  if (isNight === true) {
    return totalFare * 1.2;
  } else {
    return totalFare;
  }
}

console.log(getCngFare(2));
