// Q2: Bangladesh Weekend Machine

/**
 - fucntion that take parameter
 - switch case
 - use .toUpperCase()
 */

function getDayType(day) {
  const targetDay = day.toLowerCase();

  switch (true) {
    case targetDay === "friday" || targetDay === "saturday":
      return "Weekend";
    case targetDay === "sunday" ||
      targetDay === "monday" ||
      targetDay === "tuesday" ||
      targetDay === "wednesday" ||
      targetDay === "thursday":
      return "Working Day";
    default:
      return "Invalid Day";
  }
}

console.log(getDayType("saturday"));
