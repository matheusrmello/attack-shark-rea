import "@testing-library/jest-dom/vitest";
import { readFileSync } from "node:fs";
import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, describe, expect, it } from "vitest";
import { TopBar } from "./TopBar";

afterEach(() => { cleanup(); });

describe("TopBar", () => {
  it("keeps the brand and child content without in-app window controls", () => {
    render(<TopBar><span>Connection status</span></TopBar>);
    expect(screen.getByText("ATTACK SHARK")).toBeInTheDocument();
    expect(screen.getByText("Connection status")).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: /minimise|close window/i })).not.toBeInTheDocument();
  });

  it("allows long children to scroll without adding window controls", () => {
    const css = readFileSync("src/styles/layout.css", "utf8");
    expect(css).toMatch(/\.top-bar \.brand\s*\{[^}]*min-width:\s*0;[^}]*overflow:\s*hidden;/s);
    expect(css).toMatch(/\.top-bar-actions\s*\{[^}]*min-width:\s*0;/s);
    expect(css).toMatch(/\.top-bar-children\s*\{[^}]*min-width:\s*0;[^}]*overflow-x:\s*auto;/s);
    expect(css).not.toMatch(/\.window-controls|\.window-control/);
  });
});
