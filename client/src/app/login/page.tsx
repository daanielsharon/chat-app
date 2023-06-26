"use client";
import httpRequest from "@/helper/axios";
import { AuthContext, UserInfo } from "@/modules/auth_provider";
import { useRouter } from "next/navigation";
import React, { useContext, useEffect, useState } from "react";

const index = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { push } = useRouter();
  const { authenticated } = useContext(AuthContext);

  useEffect(() => {
    if (authenticated) {
      push("/");
      return;
    }
  }, [authenticated]);

  const handleSubmit = async (e: React.SyntheticEvent) => {
    e.preventDefault();
    try {
      const response = await httpRequest.post("/login", {
        email,
        password,
      });

      if (response.data) {
        const user: UserInfo = {
          username: response.data.username,
          id: response.data.id,
        };

        localStorage.setItem("user_info", JSON.stringify(user));

        return push("/");
      }
    } catch (error) {
      console.log("error", error);
    }
  };

  return (
    <div className="flex items-center justify-center min-w-full min-h-screen">
      <form className="flex flex-col md:w-1/5">
        <div className="text-3xl font-bold text-center">
          <span className="text-blue">Welcome!</span>
        </div>
        <input
          type="email"
          placeholder="Enter your email"
          className="p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Enter your password"
          className="p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button
          className="p-3 mt-6 rounded-md bg-blue font-bold text-white"
          type="submit"
          onClick={handleSubmit}
        >
          Login
        </button>
      </form>
    </div>
  );
};

export default index;
