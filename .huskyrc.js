/*
 * Copyright (c) 2020-present Sven Greb <development@svengreb.de>
 * This source code is licensed under the MIT license found in the LICENSE file.
 */

/**
 * The configuration for husky.
 *
 * @see https://github.com/typicode/husky
 */
module.exports = {
  hooks: {
    "pre-commit": "lint-staged",
  },
};
