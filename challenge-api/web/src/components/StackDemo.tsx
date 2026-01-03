import { motion } from "framer-motion";
import { useState } from "react";
import { z } from "zod";

const schema = z.object({
  email: z.string().email(),
});

export function StackDemo() {
  const [error, setError] = useState<string | null>(null);

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const result = schema.safeParse({ email: formData.get("email") });

    if (!result.success) {
      setError(result.error.issues[0]?.message ?? "Invalid input");
    } else {
      setError(null);
      alert("OK");
    }
  }

  return (
    <div className="space-y-6 max-w-md mx-auto">
      <motion.div
        className="w-24 h-24 rounded-2xl bg-sky-500"
        initial={{ opacity: 0, y: 24 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
      />

      <form onSubmit={handleSubmit} className="space-y-2">
        <input
          name="email"
          className="w-full rounded border border-slate-600 bg-slate-900 px-3 py-2 text-sm text-slate-100 placeholder:text-slate-500 focus:outline-none focus:ring-2 focus:ring-emerald-500"
          placeholder="输入邮箱，Zod 校验"
        />
        {error && <p className="text-sm text-red-400">{error}</p>}
        <button
          type="submit"
          className="inline-flex items-center justify-center rounded bg-emerald-500 px-4 py-2 text-sm font-medium text-white hover:bg-emerald-600"
        >
          提交
        </button>
      </form>
    </div>
  );
}
