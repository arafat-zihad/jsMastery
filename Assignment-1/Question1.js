// Q1: Value Detective

function describeValue(str) {
  const t = typeof str;

  if (str) {
    return `${t} | truthy`;
  } else {
    // tf = "falsy";
    return `${t} | falsy`;
  }
}

console.log(describeValue(225));
