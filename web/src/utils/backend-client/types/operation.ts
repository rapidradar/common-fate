/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * Common Fate
 * Common Fate API
 * OpenAPI spec version: 1.0
 */
import type { ResourceFilterOperationTypeEnum } from './resourceFilterOperationTypeEnum';

export interface Operation {
  operationType: ResourceFilterOperationTypeEnum;
  value?: string;
  attribute: string;
  values?: string[];
  operations?: Operation;
}