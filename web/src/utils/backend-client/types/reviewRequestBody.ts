/**
 * Generated by orval v6.7.1 🍺
 * Do not edit manually.
 * Common Fate
 * Common Fate API
 * OpenAPI spec version: 1.0
 */
import type { ReviewDecision } from './reviewDecision';
import type { RequestAccessGroupTiming } from './requestAccessGroupTiming';

export type ReviewRequestBody = {
  decision: ReviewDecision;
  comment?: string;
  overrideTiming?: RequestAccessGroupTiming;
};
