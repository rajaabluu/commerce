import { useState } from "react";

export default function useModal() {
  const [isShow, setIsShow] = useState(false);
  const toggle = () => setIsShow(!isShow);
  return { toggle, isShow };
}
