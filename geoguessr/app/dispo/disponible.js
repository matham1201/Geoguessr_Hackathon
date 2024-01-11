function disponible() {
  const randomBoolean = () => Math.random() < 0.5;
  return randomBoolean();
}

export default disponible;