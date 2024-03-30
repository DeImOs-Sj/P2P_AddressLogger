// This file is part of Substrate.

// Copyright (C) 2022 Parity Technologies (UK) Ltd.
// SPDX-License-Identifier: Apache-2.0

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//! Autogenerated weights for `pallet_tips`
//!
//! THIS FILE WAS AUTO-GENERATED USING THE SUBSTRATE BENCHMARK CLI VERSION 4.0.0-dev
//! DATE: 2023-09-13, STEPS: `50`, REPEAT: `20`, LOW RANGE: `[]`, HIGH RANGE: `[]`
//! WORST CASE MAP SIZE: `1000000`
//! HOSTNAME: `ip-172-31-12-189`, CPU: `Intel(R) Xeon(R) Platinum 8175M CPU @ 2.50GHz`
//! EXECUTION: ``, WASM-EXECUTION: `Compiled`, CHAIN: `Some("dev")`, DB CACHE: 1024

// Executed Command:
// ./target/release/data-avail
// benchmark
// pallet
// --chain=dev
// --steps=50
// --repeat=20
// --pallet=pallet_tips
// --extra
// --extrinsic=*
// --heap-pages=4096
// --header=./HEADER-APACHE2
// --log=warn
// --output
// ./output/pallet_tips.rs

#![cfg_attr(rustfmt, rustfmt_skip)]
#![allow(unused_parens)]
#![allow(unused_imports)]
#![allow(missing_docs)]

use frame_support::{traits::Get, weights::Weight};
use core::marker::PhantomData;

/// Weight functions for `pallet_tips`.
pub struct WeightInfo<T>(PhantomData<T>);
impl<T: frame_system::Config> pallet_tips::WeightInfo for WeightInfo<T> {
	/// Storage: `Tips::Reasons` (r:1 w:1)
	/// Proof: `Tips::Reasons` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// Storage: `Tips::Tips` (r:1 w:1)
	/// Proof: `Tips::Tips` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// The range of component `r` is `[0, 16384]`.
	fn report_awesome(r: u32, ) -> Weight {
		// Proof Size summary in bytes:
		//  Measured:  `4`
		//  Estimated: `3469`
		// Minimum execution time: 60_120_000 picoseconds.
		Weight::from_parts(65_001_713, 0)
			.saturating_add(Weight::from_parts(0, 3469))
			// Standard Error: 43
			.saturating_add(Weight::from_parts(1_983, 0).saturating_mul(r.into()))
			.saturating_add(T::DbWeight::get().reads(2))
			.saturating_add(T::DbWeight::get().writes(2))
	}
	/// Storage: `Tips::Tips` (r:1 w:1)
	/// Proof: `Tips::Tips` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// Storage: `Tips::Reasons` (r:0 w:1)
	/// Proof: `Tips::Reasons` (`max_values`: None, `max_size`: None, mode: `Measured`)
	fn retract_tip() -> Weight {
		// Proof Size summary in bytes:
		//  Measured:  `221`
		//  Estimated: `3686`
		// Minimum execution time: 57_331_000 picoseconds.
		Weight::from_parts(58_262_000, 0)
			.saturating_add(Weight::from_parts(0, 3686))
			.saturating_add(T::DbWeight::get().reads(1))
			.saturating_add(T::DbWeight::get().writes(2))
	}
	/// Storage: `Elections::Members` (r:1 w:0)
	/// Proof: `Elections::Members` (`max_values`: Some(1), `max_size`: None, mode: `Measured`)
	/// Storage: `Tips::Reasons` (r:1 w:1)
	/// Proof: `Tips::Reasons` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// Storage: `Tips::Tips` (r:0 w:1)
	/// Proof: `Tips::Tips` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// The range of component `r` is `[0, 16384]`.
	/// The range of component `t` is `[1, 3]`.
	fn tip_new(r: u32, t: u32, ) -> Weight {
		// Proof Size summary in bytes:
		//  Measured:  `367 + t * (64 ±0)`
		//  Estimated: `3832 + t * (64 ±0)`
		// Minimum execution time: 36_758_000 picoseconds.
		Weight::from_parts(36_334_070, 0)
			.saturating_add(Weight::from_parts(0, 3832))
			// Standard Error: 9
			.saturating_add(Weight::from_parts(1_971, 0).saturating_mul(r.into()))
			// Standard Error: 57_709
			.saturating_add(Weight::from_parts(1_337_520, 0).saturating_mul(t.into()))
			.saturating_add(T::DbWeight::get().reads(2))
			.saturating_add(T::DbWeight::get().writes(2))
			.saturating_add(Weight::from_parts(0, 64).saturating_mul(t.into()))
	}
	/// Storage: `Elections::Members` (r:1 w:0)
	/// Proof: `Elections::Members` (`max_values`: Some(1), `max_size`: None, mode: `Measured`)
	/// Storage: `Tips::Tips` (r:1 w:1)
	/// Proof: `Tips::Tips` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// The range of component `t` is `[1, 3]`.
	fn tip(t: u32, ) -> Weight {
		// Proof Size summary in bytes:
		//  Measured:  `588 + t * (112 ±0)`
		//  Estimated: `4053 + t * (112 ±0)`
		// Minimum execution time: 28_593_000 picoseconds.
		Weight::from_parts(30_467_919, 0)
			.saturating_add(Weight::from_parts(0, 4053))
			// Standard Error: 31_291
			.saturating_add(Weight::from_parts(73_851, 0).saturating_mul(t.into()))
			.saturating_add(T::DbWeight::get().reads(2))
			.saturating_add(T::DbWeight::get().writes(1))
			.saturating_add(Weight::from_parts(0, 112).saturating_mul(t.into()))
	}
	/// Storage: `Tips::Tips` (r:1 w:1)
	/// Proof: `Tips::Tips` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// Storage: `Elections::Members` (r:1 w:0)
	/// Proof: `Elections::Members` (`max_values`: Some(1), `max_size`: None, mode: `Measured`)
	/// Storage: `System::Account` (r:1 w:1)
	/// Proof: `System::Account` (`max_values`: None, `max_size`: Some(128), added: 2603, mode: `MaxEncodedLen`)
	/// Storage: `Tips::Reasons` (r:0 w:1)
	/// Proof: `Tips::Reasons` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// The range of component `t` is `[1, 3]`.
	fn close_tip(t: u32, ) -> Weight {
		// Proof Size summary in bytes:
		//  Measured:  `692 + t * (60 ±0)`
		//  Estimated: `4143 + t * (71 ±1)`
		// Minimum execution time: 117_623_000 picoseconds.
		Weight::from_parts(115_579_473, 0)
			.saturating_add(Weight::from_parts(0, 4143))
			// Standard Error: 1_216_708
			.saturating_add(Weight::from_parts(10_801_842, 0).saturating_mul(t.into()))
			.saturating_add(T::DbWeight::get().reads(3))
			.saturating_add(T::DbWeight::get().writes(3))
			.saturating_add(Weight::from_parts(0, 71).saturating_mul(t.into()))
	}
	/// Storage: `Tips::Tips` (r:1 w:1)
	/// Proof: `Tips::Tips` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// Storage: `Tips::Reasons` (r:0 w:1)
	/// Proof: `Tips::Reasons` (`max_values`: None, `max_size`: None, mode: `Measured`)
	/// The range of component `t` is `[1, 3]`.
	fn slash_tip(t: u32, ) -> Weight {
		// Proof Size summary in bytes:
		//  Measured:  `269`
		//  Estimated: `3734`
		// Minimum execution time: 26_915_000 picoseconds.
		Weight::from_parts(29_724_450, 0)
			.saturating_add(Weight::from_parts(0, 3734))
			// Standard Error: 97_138
			.saturating_add(Weight::from_parts(47_331, 0).saturating_mul(t.into()))
			.saturating_add(T::DbWeight::get().reads(1))
			.saturating_add(T::DbWeight::get().writes(2))
	}
}