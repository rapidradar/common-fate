/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * Common Fate
 * Common Fate API
 * OpenAPI spec version: 1.0
 */
import type { CreateAccessRuleTargetFieldFilterExpessions } from './createAccessRuleTargetFieldFilterExpessions';

/**
 * a request body for creating a Access Rule Target
 */
export interface CreateAccessRuleTarget {
  targetGroupId: string;
  fieldFilterExpessions: CreateAccessRuleTargetFieldFilterExpessions;
}
