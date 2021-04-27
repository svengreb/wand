/*
 * Copyright (c) 2019-present Sven Greb <development@svengreb.de>
 * This source code is licensed under the MIT license found in the LICENSE file.
 */

/**
 * Configurations for lint-staged.
 *
 * @see https://github.com/okonet/lint-staged#configuration
 */
module.exports = {
  "*.{css,html,js,json,yaml,yml}": "prettier --check",
  "*.md": ["remark --no-stdout", "prettier --check"],
};
