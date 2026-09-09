export type SessionMergeId = string;

export interface SessionMergeEvent {
  id: SessionMergeId;
  version: number;
  occurredAt: string;
  source: string;
}

export const STORAGE_KIND = "IndexedDB" as const;
export const FRAMEWORK_KIND = "React" as const;
