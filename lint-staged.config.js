/*
 * Copyright (c) 2020-present Sven Greb <development@svengreb.de>
 * This source code is licensed under the MIT license found in the LICENSE file.
 */

/**
 * The configuration for lint-staged.
 *
 * @see https://github.com/okonet/lint-staged#configuration
 */
module.exports = {
  "*": "prettier --check",
  "*.md": "remark --no-stdout",
};
