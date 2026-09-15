// Q3: Username Gatekeeper

function validateUsername(value) {
  const str = value.toLowerCase();
  const strLength = str.length;
  // const reserved = str.includes("admiin")

  if (strLength < 4) {
    return "Too Short";
  } else if (str.includes(" ")) {
    return "No Space Allowed";
  } else if (str.includes("admin")) {
    return "Reserved Word";
  } else {
    return "Available";
  }
}

console.log(validateUsername("kapani datta"));
