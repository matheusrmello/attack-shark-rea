import type { ReactNode } from "react";

export function TopBar({
  title = "Mouse configuration",
  subtitle = "Attack Shark X6",
  children,
}: {
  title?: string;
  subtitle?: string;
  children?: ReactNode;
}) {
  return (
    <header className="titlebar top-bar">
      <div className="appname brand">
        <img className="app-icon" src="/attack-shark-x6.png" alt="" />
        <span className="brand-wordmark">ATTACK SHARK</span>
        <small>{subtitle || "X6"}</small>
        <span className="sr-only">{title}</span>
      </div>
      <div className="top-bar-actions">
        <div className="top-bar-children">{children}</div>
      </div>
    </header>
  );
}
