import assert from "node:assert/strict";
import test from "node:test";
import { FRAMEWORK_KIND, STORAGE_KIND, type SessionMergeEvent } from "../src/domain/contracts.js";

test("领域事件保留版本和来源", () => {
  const event: SessionMergeEvent = { id: "sample", version: 1, occurredAt: "2026-09-09T00:00:00Z", source: "fixture" };
  assert.equal(event.version, 1);
  assert.equal(FRAMEWORK_KIND, "React");
  assert.equal(STORAGE_KIND, "IndexedDB");
});
